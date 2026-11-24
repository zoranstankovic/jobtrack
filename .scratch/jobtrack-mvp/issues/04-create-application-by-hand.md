# 04: Create an Application by hand

**What to build:** The user opens Applications and creates an Application for an existing Company. The fields are: position title, Channel (Agentur für Arbeit, LinkedIn, Indeed, instaffo, JOIN, Direct, Other), link, Location, Salary Expectation (yearly gross EUR), Job Description and Notes. The user also chooses the initial Application Status and its date (default Applied, today). Notes start prefilled with the template headings (Summary, Form Answers, Strengths, Weak Spots, Interview Prep, Documents). Saving stores the Application together with its first Status Change. The Applications list shows Company, position, Channel, current Application Status and the date of the last change. The detail page shows every field: Notes rendered from Markdown, the Job Description as plain text with its line breaks kept, and the Salary Expectation as yearly and monthly (÷12).

**Teaches:** foreign keys and relations, enums as constrained text, a transaction across two inserts, rendering Markdown safely with `goldmark` (raw HTML disabled), computed display values.

**Blocked by:** 02

**Status:** ready-for-agent

- [ ] The create form requires Company (choose from existing), position title and Channel; Salary Expectation, if given, is a whole number that is not negative
- [ ] Notes are prefilled with the template headings
- [ ] The initial Application Status can be any of the 9 statuses, with a date that may be in the past; the Application and its first Status Change are saved in one transaction
- [ ] The Applications list shows Company, position, Channel, current Application Status and the date of the last change, most recently changed first
- [ ] The current Application Status is defined as the Status of the latest Status Change by date (latest recorded wins on the same date)
- [ ] The detail page renders Notes as Markdown (raw HTML not rendered) and the Job Description with its line breaks kept
- [ ] The detail page shows the Salary Expectation per year and per month
- [ ] Tests through the HTTP seam cover: create with a past date, validation errors, list content, detail rendering
