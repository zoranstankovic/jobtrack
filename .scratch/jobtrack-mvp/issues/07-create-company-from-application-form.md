# 07: Create a Company from the Application form

**What to build:** While creating or editing an Application, the user sees that the Company is missing. Next to the Company picker, "New Company" opens a small inline form (name, website, ATS, Careers URL). Saving it creates the Company and selects it in the Application form. What the user already typed in the Application form is kept.

**Teaches:** nested HTMX interactions (a form inside a form's flow), targeting a different element than the one that triggered the request, out-of-band swaps (`hx-swap-oob`), reusing the Company validation.

**Blocked by:** 04

**Status:** ready-for-agent

- [ ] "New Company" in the Application form opens an inline Company form without leaving the page
- [ ] Saving creates the Company (same validation as 02, including unique name ignoring case) and selects it in the Company picker
- [ ] Values already entered in the Application form are kept
- [ ] Cancel closes the inline form without changes
- [ ] Tests through the HTTP seam cover: create a Company inline, then create an Application with it; the duplicate-name error inline
