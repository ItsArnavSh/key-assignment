# key-assignment
# How it works
- Its for a simulated environment, so we generate a sample codebase with thousands of files and code using osdc/resrap package
- Now we take all the keys present in the inventory.json and convert to a DS to perform  AhoCorasick algorithm, which is efficient when we have to look for multiple strings in a single codebase
- If any matches are found, the report starts being generated
- We take the url(In this code we are using sample URL since its supposed to be a simulated exercise, but in real world scenario, we would download and scan repositories, so we would take that url) and we are scanning its repo for its authors, all info about them, potential links, personal websites, socials etc all that can be deducted from their github account, is compiled and a detailed report is sent over to the slack channel
