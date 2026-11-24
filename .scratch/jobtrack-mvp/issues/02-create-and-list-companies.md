# 02: Create and list Companies

**What to build:** The user opens Companies, fills in a form (name, website, ATS from the fixed list with default Unknown, Careers URL, notes) and saves. The new Company appears in the list without a full page reload. The list shows each Company's name, ATS and Careers URL. Invalid input is shown next to the fields, and what the user typed is kept. A second Company with the same name (ignoring letter case) is rejected with a clear message.

**Teaches:** your first migration and table, writing SQL queries for `sqlc`, an Echo handler that renders a full page or a fragment depending on the `HX-Request` header, HTMX form submission, server-side validation.

**Blocked by:** 01

**Status:** ready-for-agent

- [ ] Companies page lists all Companies, sorted by name
- [ ] The create form accepts name (required), website, ATS (Greenhouse, Lever, Ashby, Workday, Personio, SmartRecruiters, Other, Unknown; default Unknown), Careers URL and notes
- [ ] Submitting a valid form adds the Company to the list through an HTMX swap, without a full reload
- [ ] A missing name shows an error next to the field and keeps the other entered values
- [ ] A name that matches an existing Company ignoring case (e.g. "zalando" vs "Zalando") is rejected with a message
- [ ] The Careers URL is shown as a clickable link that opens in a new tab
- [ ] Tests through the HTTP seam cover: create, list, validation error, duplicate name
