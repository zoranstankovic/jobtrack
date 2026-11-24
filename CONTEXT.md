# Jobtrack

A personal tool for one person: it tracks the jobs they have applied to and collects job postings from public sources, based on their saved searches.

## Language

### Job hunting

**Company**:
An employer the user is following or has applied to. Every Application belongs to exactly one Company.
_Avoid_: Firma, Employer, Arbeitgeber, Organization

**ATS**:
The Applicant Tracking System a Company uses for job postings and applications. It is one value from a fixed list: Greenhouse, Lever, Ashby, Workday, Personio, SmartRecruiters, Other, Unknown.
_Avoid_: Recruiting platform, HR system

**Careers URL**:
The address of a Company's main page that lists all its open positions, usually hosted on its ATS.
_Avoid_: Job board, Karriereseite

**Application**:
The user's act of applying for a job, along with what happens to it afterward. It always belongs to one Company and is sent through one Channel. It can exist without a Job Posting (for example, an application sent through LinkedIn and entered by hand).
_Avoid_: Prijava, Bewerbung, Job

**Application Status**:
How far an Application has progressed. The values are: Planned, Applied, Technical Assessment, Interview, Offer, Accepted, Rejected, Withdrawn (the user gave up), Ghosted (no answer). Planned means "I'm interested and plan to apply". Any status may follow any other.
_Avoid_: Stage, Phase, State

**Status Change**:
A single entry in an Application's history that records when it moved to which Application Status. Together these entries form the Application's timeline.
_Avoid_: Event, Log entry, Transition

**Channel**:
Where the user sent an Application. It is one value from a fixed list: Agentur für Arbeit, LinkedIn, Indeed, instaffo, JOIN, Direct, Other. It is different from a Source (where jobtrack collects Job Postings from) and from an ATS (the system the Company uses).
_Avoid_: Platform, Kanal, Applied via

**Job Description**:
The full text of a job advertisement, kept as a copy on the Application: exactly what the user applied to, even if the original later changes or disappears.
_Avoid_: Posting text, Oglas, Stellenbeschreibung

**Salary Expectation**:
The yearly gross amount in EUR that the user stated in the Application.
_Avoid_: Salary, Plata, Gehaltsvorstellung

### Collecting postings

**Job Posting**:
A job advertisement published by an employer and collected from a Source. It is uniquely identified by its Source and that Source's own identifier, and it has at most one Application. The user only sees it; it is not theirs.
_Avoid_: Job, Posao, Oglas, Stellenangebot, Listing

**Posting Status**:
The user's triage of a Job Posting: New (not yet reviewed) or Dismissed (not interested, do not show again). Whether the user applied is not a Posting Status: it follows from the Application linked to the Job Posting.
_Avoid_: Bookmark, Saved, Read/unread

**Search Profile**:
A saved set of criteria used to collect Job Postings: keywords, location, radius, working time and how many days back to look. One Job Posting can be found by several Search Profiles.
_Avoid_: Profil, Filter, Query, Saved search

**Source**:
An outside system that publishes Job Postings. For now this is only the Bundesagentur für Arbeit.
_Avoid_: Provider, Feed, Integration

**Refresh**:
Collecting new Job Postings from a Source using the Search Profiles, started by the user.
_Avoid_: Sync, Import, Scrape, Fetch
