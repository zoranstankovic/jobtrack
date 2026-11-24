# 12: Match Companies for Job Postings

**What to build:** When the user creates an Application from a Job Posting, jobtrack proposes the Company. It first looks for a Company with the same BA employer identifier (`arbeitgeberKundennummerHash`), then for one with the same name ignoring case. When a match is found, it is preselected and the user can change it. When there is no match, a new Company is proposed with the posting's employer name, and creating it stores the BA employer identifier so the next posting from that employer matches automatically, even under a slightly different name. When an existing Company is chosen for a posting and that Company has no BA employer identifier yet, the identifier is remembered.

**Teaches:** matching logic with an order of preference, queries over several columns, keeping a module's rules (Tracking) separate from the Source details, test-first development for edge cases.

**Blocked by:** 11

**Status:** ready-for-agent

- [ ] Written test-first (TDD)
- [ ] A Company with the same BA employer identifier is preselected, even if its name differs (e.g. "Zalando" vs "Zalando SE")
- [ ] Otherwise a Company with the same name ignoring case is preselected
- [ ] Otherwise a new Company with the posting's employer name is proposed; creating it stores the BA employer identifier
- [ ] The user can override the proposed Company with any other or a new one
- [ ] Choosing an existing Company without a BA employer identifier for a posting stores the identifier on that Company
