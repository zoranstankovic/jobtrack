# 08: Search Profiles CRUD

**What to build:** The user opens Search Profiles and creates, edits and deletes them. A Search Profile has a name, keywords, location, radius in km, working time (any combination of full-time, part-time, remote; none selected means any) and days back (0–100). The list shows every profile with its criteria in short form.

**Teaches:** storing a set of values (Postgres arrays or a join table), checkbox groups in HTML forms and how Echo binds them, range validation.

**Blocked by:** 01

**Status:** ready-for-agent

- [ ] The create/edit form requires name and keywords; radius is a whole number that is not negative; days back is 0–100
- [ ] Working time is a set of full-time, part-time and remote, and saving with none selected is allowed (meaning any)
- [ ] The list shows each Search Profile's name and a short summary of its criteria
- [ ] Edit and delete work inline with HTMX; delete asks for confirmation
- [ ] Tests through the HTTP seam cover: create, validation errors (days back 101, missing keywords), edit, delete
