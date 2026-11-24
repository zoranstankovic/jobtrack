# 11: Create an Application from a Job Posting

**What to build:** On a New Job Posting the user clicks "Create Application". The Application form opens prefilled from the posting: position title, Location, link to the original posting, Channel = Agentur für Arbeit, and Notes with the template. The full Job Description is fetched from BA (`GET /pc/v4/jobdetails/{base64(referenznummer)}`, field `stellenangebotsBeschreibung`) when the form opens and copied into the form. If that fetch fails, the form still opens with a notice and an empty Job Description. The user picks or creates the Company (with the inline form from 07, the name prefilled with the posting's employer name) and saves. The Application is linked to the posting, and the posting drops out of the New postings list. A Job Posting has at most one Application: trying again leads the user to the existing one.

**Teaches:** a second Source endpoint and base64 encoding, prefilling forms from another entity, `UNIQUE` constraints as a domain rule, redirects in an HTMX flow (`HX-Redirect`), designing for partial failure.

**Blocked by:** 07, 10

**Status:** ready-for-agent

- [ ] Written test-first (TDD) for the Job Description copy, including the failure path
- [ ] "Create Application" opens the form prefilled with title, Location, link, Channel = Agentur für Arbeit and the Notes template
- [ ] The Job Description is fetched from the fake BA details endpoint (the identifier sent is the base64 of the `referenznummer`) and appears in the form
- [ ] When the details fetch fails, the form opens anyway with a visible notice and an empty Job Description
- [ ] The inline "New Company" form is prefilled with the posting's employer name
- [ ] The saved Application is linked to the Job Posting; the posting no longer appears among New postings
- [ ] A second "Create Application" on the same posting goes to the existing Application instead of creating another one
- [ ] The Job Description stays on the Application unchanged regardless of what happens to the posting
