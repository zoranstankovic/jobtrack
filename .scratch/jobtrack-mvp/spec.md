# Spec: jobtrack MVP

Status: ready-for-agent

## Problem Statement

I'm looking for a job in Germany and I track my applications in a Notion database that I took from an online template. It is generic. It does not show me how long I've been waiting on each Company, and it does not connect to the places where postings are published. For each Application I copy the Job Description, my Salary Expectation, my form answers and my interview notes into it by hand.

I also search the Bundesagentur für Arbeit job board by hand, over and over with the same criteria. I see the same Job Postings again and again, and I can't hide the ones I've already rejected.

I want one place that tracks every Application from the moment I plan it until it ends, and that collects new Job Postings matching my saved criteria so I only look at what's new.

This is also a learning project. I'm new to Go, Echo, HTMX, Postgres and Docker, and building it should teach me all of them.

## Solution

jobtrack is a single-user web application that starts with one `docker compose up`. It has two halves that are connected to each other.

1. **Tracking.** I record the Companies I care about (including their ATS and Careers URL) and my Applications to them. Each Application has a timeline of Status Changes, a Channel, a Salary Expectation, a copy of the Job Description, and Markdown Notes that start from my usual template.
2. **Collecting.** I save Search Profiles (keywords, location, radius, working time, how many days back to look). A Refresh pulls matching Job Postings from the Bundesagentur für Arbeit (the only Source for now). Each Job Posting is stored once, and I triage it: I Dismiss it, or I create an Application from it. Creating an Application from a posting copies its Job Description and prefills everything it can.

The Dashboard shows new Job Postings and my active Applications, including how many days have passed since each one last changed.

The UI is HTML rendered on the server and swapped into the page by HTMX. There is no JSON or GraphQL API (ADR 0001).

## User Stories

### Running the application

1. As the user, I want to start the whole application (app and database) with one command, so that running it takes no setup knowledge.
2. As the user, I want database migrations to run automatically when the app starts, so that I never have to run them by hand.
3. As the user, I want my data to survive restarts of the containers, so that I don't lose my Applications.
4. As the user, I want to configure the database connection, the listening port and the Source's address through environment variables, so that the same image runs anywhere.

### Companies

5. As the user, I want to create a Company with a name, website, ATS, Careers URL and notes, so that I have one record per employer I care about.
6. As the user, I want to pick the ATS from a fixed list (Greenhouse, Lever, Ashby, Workday, Personio, SmartRecruiters, Other, Unknown), so that the data stays consistent and can later drive ATS integrations.
7. As the user, I want ATS to default to Unknown, so that I can create a Company without researching it first.
8. As the user, I want to see a list of all Companies with their ATS and number of Applications, so that I know who I've dealt with.
9. As the user, I want to edit a Company in place without a full page reload, so that small corrections are quick.
10. As the user, I want to be stopped from creating two Companies with the same name (ignoring letter case), so that I don't end up with duplicates.
11. As the user, I want to open a Company and see all its Applications, so that I remember my history with that employer before I apply again.
12. As the user, I want to click a Company's Careers URL, so that I can jump straight to its open positions.
13. As the user, I want to be prevented from deleting a Company that still has Applications, so that I never orphan my history.
14. As the user, I want to delete a Company that has no Applications, so that I can clean up mistakes.

### Applications

15. As the user, I want to create an Application by hand for a Company, with a position title, Channel, link, Location, Salary Expectation, Job Description and Notes, so that I can track applications I found outside jobtrack.
16. As the user, I want to create a new Company directly from the Application form when it doesn't exist yet, so that I don't have to leave the form.
17. As the user, I want to pick the Channel from a fixed list (Agentur für Arbeit, LinkedIn, Indeed, instaffo, JOIN, Direct, Other), so that I can see later where I applied.
18. As the user, I want to choose the initial Application Status (usually Planned or Applied) and its date when I create an Application, so that an application I enter late gets the right date.
19. As the user, I want the Notes of a new Application to be prefilled with my template headings (Summary, Form Answers, Strengths, Weak Spots, Interview Prep, Documents), so that every Application has the same structure.
20. As the user, I want Notes written in Markdown and shown formatted, so that I can use headings, lists and links (for example to my CV and cover letter).
21. As the user, I want to paste the full Job Description into the Application, so that I keep exactly what I applied to.
22. As the user, I want the Job Description shown as readable text with its line breaks kept, so that long postings stay readable.
23. As the user, I want to enter my Salary Expectation once as a yearly gross amount in EUR and also see the monthly amount (divided by 12), so that I don't keep two numbers in sync.
24. As the user, I want to see a list of all Applications with Company, position, Channel, current Application Status and the date of the last change, so that I have an overview.
25. As the user, I want to filter the Application list to only active ones (not Accepted, Rejected, Withdrawn or Ghosted), so that I can focus on what is still open.
26. As the user, I want to open an Application and see all its fields, its Notes and its timeline, so that I can prepare for a call with that Company.
27. As the user, I want to edit an Application's fields, so that I can fill in details later.
28. As the user, I want to delete an Application (after confirming), so that I can remove ones I entered by mistake.

### Application Status and timeline

29. As the user, I want to change an Application's Application Status from the list or detail view without a page reload, so that updating status takes one click.
30. As the user, I want every status change recorded as a Status Change with its date, so that I have a full timeline.
31. As the user, I want to give a Status Change a date in the past, so that I can record something that happened a few days ago.
32. As the user, I want to move from any Application Status to any other (for example Ghosted → Interview), so that the tool never blocks what really happened.
33. As the user, I want to record the same Application Status twice (for example two Interview rounds), so that every round shows up on the timeline.
34. As the user, I want the current Application Status to be the one from the latest Status Change by date, so that entering a past event doesn't overwrite a newer one.
35. As the user, I want the date I applied derived from the timeline rather than entered separately, so that the two can never contradict each other.
36. As the user, I want to delete a wrong Status Change (as long as at least one remains), so that I can fix mistakes in the timeline.

### Search Profiles

37. As the user, I want to create a Search Profile with a name, keywords, location, radius in km, working time (any combination of full-time, part-time and remote) and how many days back to look, so that I can save the searches I'd otherwise repeat by hand.
38. As the user, I want "days back" limited to 0–100, so that I can't save a search the Source would reject.
39. As the user, I want to list, edit and delete Search Profiles, so that I can adjust my criteria as the search goes on.
40. As the user, I want deleting a Search Profile to keep the Job Postings it already found, so that I don't lose postings I'm still deciding about.
41. As the user, I want to see how many New Job Postings each Search Profile has, so that I know which searches are productive.

### Refresh and Job Postings

42. As the user, I want a "Refresh all" button on the Dashboard that runs every Search Profile against the Source, so that I collect new Job Postings in one click.
43. As the user, I want a Refresh button on each Search Profile, so that I can check one search right after creating or editing it.
44. As the user, I want each Refresh to fetch at most the 100 most recent Job Postings per Search Profile, so that a broad search can't flood the app.
45. As the user, I want a Job Posting stored only once even when several Search Profiles or several Refreshes find it, so that I never see duplicates.
46. As the user, I want to see which Search Profiles found each Job Posting, so that I understand why it appeared.
47. As the user, I want a Refresh to never change my triage of Job Postings that already exist, so that Dismissed postings stay dismissed.
48. As the user, I want to see after a Refresh how many Job Postings were new, so that I know whether it was worth it.
49. As the user, I want a clear error message when the Source is unreachable or answers with an error, and nothing lost, so that a failed Refresh is harmless.
50. As the user, I want each Job Posting to show title, employer name, location, publication date and a link to the original posting on the Source's website, so that I can judge it quickly.
51. As the user, I want to Dismiss a Job Posting and have it disappear from the list immediately, so that triage is fast.
52. As the user, I want Dismissed Job Postings never shown again, even if a later Refresh finds them again, so that I make each decision only once.
53. As the user, I want Job Postings that already have an Application to drop out of the New list, so that the list only shows undecided postings.

### Creating an Application from a Job Posting

54. As the user, I want a "Create Application" button on each New Job Posting, so that turning a posting into an Application takes one step.
55. As the user, I want the Application form prefilled from the Job Posting (position title, Location, link, Channel = Agentur für Arbeit), so that I don't retype anything.
56. As the user, I want the full Job Description fetched from the Source and copied into the form, so that my Application keeps it even after the posting is taken down.
57. As the user, I want the form to open even if fetching the Job Description fails, showing a notice and an empty Job Description I can paste into, so that the Source never blocks me.
58. As the user, I want jobtrack to match the posting's employer to an existing Company, first by the Source's employer identifier and then by name ignoring case, so that I don't create duplicate Companies.
59. As the user, I want a new Company proposed with the posting's employer name (and the Source's employer identifier remembered) when no match exists, so that the Company is created only when I actually act.
60. As the user, I want to override the proposed Company, so that I can fix a wrong match (for example "Zalando SE" vs an existing "Zalando").
61. As the user, I want to be taken to the existing Application if I try to create a second one from the same Job Posting, so that one posting never has two Applications.
62. As the user, I want to still apply to the same Company several times for different positions or at different times, so that real job-hunting patterns are supported.

### Dashboard

63. As the user, I want the Dashboard as the home page, so that I see what needs my attention right away.
64. As the user, I want a "New postings" section with New Job Postings, newest first, each with its Search Profiles and Dismiss / Create Application buttons, so that I can triage from one place.
65. As the user, I want an "Active applications" section sorted by last change (oldest first), each showing Company, position, current Application Status and days since the last change, so that I see who needs a follow-up or should be marked Ghosted.
66. As the user, I want the Dashboard to update in place after Refresh, Dismiss or a status change, so that it never feels like a slow full-page app.

### Usability

67. As the user, I want form validation errors shown next to the fields without losing what I typed, so that fixing mistakes is painless.
68. As the user, I want simple navigation between Dashboard, Applications, Companies and Search Profiles, so that I can always get where I need to go.
69. As the user, I want the whole UI in English, matching the vocabulary in the glossary, so that the words in the UI and in the code are the same.

## Implementation Decisions

### Architecture and stack

- Go web application on Echo. HTML is rendered on the server with `templ` components, and interactivity comes from HTMX. Handlers return either full pages or fragments, based on whether the request came from HTMX (`HX-Request` header). There is no JSON or GraphQL API (ADR 0001).
- Postgres is accessed with `pgx`. Queries are written as SQL and turned into typed Go code by `sqlc`. Migrations use `goose` and run automatically at app startup.
- Markdown (Notes) is rendered with `goldmark`, with raw HTML disabled. The Job Description is shown as plain text with line breaks kept.
- All configuration comes from environment variables: database URL, listen address and the Source (BA API) base URL. The BA base URL is configurable so that tests can point it at a fake server. This is the only test seam that reaches outside the app.
- Docker: a multi-stage Dockerfile (build stage with Go and `templ`, small runtime stage) plus a compose file with the app and Postgres (healthcheck, a named volume for data, the app waits for a healthy database).
- Single user, no authentication.
- Everything in the repo (code, comments, UI, docs) is in English.

### Modules

- **Tracking**: Companies, Applications and Status Changes. It owns all rules:
  - Company names are unique ignoring case.
  - A Company with Applications cannot be deleted.
  - An Application always has at least one Status Change, and its current Application Status comes from the latest Status Change by date.
  - A Job Posting has at most one Application.
  - It provides matching a Company for a Job Posting's employer: by employer identifier from the Source, then by name ignoring case.
- **Collecting**: Search Profiles, Job Postings and Refresh.
  - Refresh runs each Search Profile through the Source.
  - It upserts Job Postings by (Source, the Source's identifier) and links each one to the Search Profile that found it.
  - It never changes an existing posting's fields or Posting Status.
  - It reports how many postings were new.
- **BA Source client**: the only code that knows the Bundesagentur für Arbeit API. Its narrow interface does two things.
  - **Search** with a Search Profile. It translates keywords → `was`, location → `wo`, radius → `umkreis`, working time → `arbeitszeit` (full-time `vz`, part-time `tz`, remote `ho`, separated by `;`), days back → `veroeffentlichtseit`, and always sends `angebotsart=1`. It fetches at most 2 pages of 50 and returns Job Postings in domain terms.
  - **Get the Job Description** for one posting by its identifier.
- **Web**: Echo routes, handlers and `templ` views (pages and fragments for every HTMX interaction). Handlers stay thin and delegate to Tracking and Collecting.

### Facts about the BA API (verified with live calls on 2026-11-24; the published docs are partly wrong)

- Base URL `https://rest.arbeitsagentur.de/jobboerse/jobsuche-service`, header `X-API-Key: jobboerse-jobsuche`.
- Search: `GET /pc/v6/jobs` (v4 returns 403).
  - Results are in `ergebnisliste`.
  - Per item: `referenznummer` (identifier), `stellenangebotsTitel` (title), `firma` (employer name), `arbeitgeberKundennummerHash` (employer identifier), `stellenlokationen[].adresse` (`ort`, `plz`, `region`), `datumErsteVeroeffentlichung`, `aenderungsdatum`.
  - Pagination: `maxErgebnisse`, `page`, `size`.
- Details: `GET /pc/v4/jobdetails/{base64(referenznummer)}` (v3 and v2 return 403). The full text is in `stellenangebotsBeschreibung`.
- Link to the original posting for people: the Source's public job detail page for the `referenznummer`.
- The search matches by occupation, not by exact keyword, so results are broad. That's why triage exists.
- No documented rate limit or terms of use.

### Data model (described in words)

- **Company**: name (unique, ignoring case), website, ATS (enum, default Unknown), Careers URL, notes, optional BA employer identifier.
- **Application**: Company (required), position title (required), Channel (enum, required), link, Location, Salary Expectation (whole EUR per year, optional, not negative), Job Description, Notes (Markdown), optional Job Posting (unique when present), created/updated timestamps.
- **Status Change**: Application, Application Status (enum of 9), date it happened, time it was recorded. When two have the same date, the one recorded later counts as more recent.
- **Search Profile**: name (required), keywords (required), location, radius km (not negative), working time (a set of full-time, part-time, remote; empty means any), days back (0–100).
- **Job Posting**: Source (enum, only BA for now), the Source's identifier (unique together with Source), title, employer name, BA employer identifier, location (text), publication date, Posting Status (New/Dismissed, default New), first-seen timestamp.
- **Search Profile ↔ Job Posting**: many-to-many link recording which profiles found which postings.
- Enums are stored as constrained text values.

### Key interactions

- **Status change**: a small HTMX form (status and date, defaulting to today). On submit, the server re-renders the Application's row or detail fragment and the timeline.
- **Dismiss**: one HTMX request. The posting's element is removed from the page.
- **Refresh all / Refresh one**: HTMX request. The New postings fragment is re-rendered, along with a short summary ("12 new postings").
- **Create Application from a Job Posting**: opens the Application form prefilled from the posting. The Job Description is fetched from the Source when the form opens. The Company is proposed by matching and can be changed or created inline. On submit, the Application and its first Status Change are stored and linked to the posting. If the posting already has an Application, the user is redirected to it.
- **Validation errors**: the server re-renders the form fragment with errors and the submitted values.

### Build order (each step is a working app)

0. Skeleton: Go module, Echo, `templ`, Postgres in compose, migrations, `sqlc`, a placeholder page, the test harness (HTTP seam + Postgres + fake BA server).
1. Companies.
2. Applications with Status Changes and the timeline.
3. Search Profiles.
4. Refresh: BA Source client, Job Posting storage, deduplication, Dismiss.
5. Dashboard and Create Application from a Job Posting (with Job Description copy and Company matching).

## Testing Decisions

- **One seam: the application's HTTP interface.**
  - Tests start the real Echo app wired to a real Postgres (from Docker) and send HTTP requests the way HTMX would: form posts, `HX-Request` headers, button actions.
  - They check the returned HTML and/or what later pages show.
  - Tests never call internal functions, never check SQL, and never mock the database.
- **The Source is faked at the network level.**
  - An `httptest` server stands in for the BA API and serves recorded JSON responses shaped like the live responses above.
  - The app gets its address through configuration.
  - Tests can also make the fake fail (5xx, timeout, bad JSON) to check error handling.
- **No fake clock.** Time-related behavior (days since last change, current status ordering) is tested by entering Status Changes with dates in the past, relative to the real current date.
- **Isolation**: each test starts from a clean database state.
- **What makes a good test**:
  - It describes behavior the user can see ("after a Refresh that returns a posting I dismissed, it still doesn't appear"), not how the code produces it.
  - It survives refactoring of handlers, queries and the client.
  - Its assertions target meaningful content (titles, statuses, counts), not exact markup.
- **TDD (test first)** for the most important logic:
  - Refresh, deduplication across profiles and Refreshes, keeping Posting Status
  - translating a Search Profile into Source query parameters (checked through the requests the fake server receives)
  - the Job Description copy, including when fetching it fails
  - Company matching
  - current Application Status ordering with past dates
- Lighter tests cover the CRUD screens from step 1 on.
- **Prior art**: none yet, since the repo is empty. Step 0 sets up the test harness, and the first Company tests in step 1 set the pattern that all later tests follow.

## Out of Scope

- Sources other than the Bundesagentur für Arbeit, including pulling postings from ATS job boards (Greenhouse, Lever, Ashby, Personio). The ATS and Careers URL are only stored for now.
- Automatic or scheduled Refresh in the background.
- Detecting Job Postings that were removed or expired at the Source.
- File uploads (CV, cover letter). Links go into Notes.
- AI-generated summaries of Job Descriptions.
- Importing existing Applications from Notion (planned later as a one-time CSV import).
- Offer amounts, contact people, interview scheduling, reminders and notifications.
- Statistics and charts.
- Authentication and multiple users.
- Publishing the Docker image to a registry, and CI.
- A JSON or GraphQL API (ADR 0001).
- Translations (i18n). The UI is English only.

## Further Notes

- The glossary is in `CONTEXT.md`. Use its terms (Application, Job Posting, Search Profile, Source, Company, ATS, Careers URL, Channel, Application Status, Status Change, Posting Status, Refresh, Job Description, Salary Expectation) everywhere: code, UI, tests, tickets.
- This is a learning project for someone new to Go, Echo, HTMX, Postgres and Docker. Working mode:
  - Claude writes the boilerplate (Docker, compose, configuration, skeleton, test harness).
  - The user writes the core (handlers, SQL queries, HTMX interactions) with guidance and review.
  - Tickets should be small enough to learn from and should name the concept each one teaches.
- The BA API is unofficial in practice. Its published spec doesn't match live behavior (see above), so the recorded test fixtures should be taken from real responses.
