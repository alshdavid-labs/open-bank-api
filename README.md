# Open Bank API

This project aims to be an open source effort to reverse engineer bank APIs.

## Why

Modern banks provide dated and feature deprived online facilities with poor quality data which, in some cases, result in costly manual handling of banking information for tasks that would be trivial with easily queryable high quality data.

Integrating directly with a bank's API allows for implementing features ourselves. 

Examples of such features are:
- The need for more complex automation tasks
- The need to automate the mirroring/archiving of financial records beyond what is offered by the banks
- Associating records locally with receipts 
- Correcting transaction dates and authorization dates

### Examples / Hall of Shame

- CommonWealth Bank in Australia only offers you data on your account for up to 7 years however records over 2 years old are only accessible as unparsable PDF files. Data on closed accounts is only available for 5 years as printed statements sent to your mail box. They cite the cost of storage as the reasoning.
- All Australian banks do not store transactions using authorization date, instead they use the post date which can be up to 5 working days delayed. This makes it difficult to ascertain when a purchase occurred for personal financial audits.

## Contributions

Please contribute!

It's the ambition of this project to contain API wrappers for as many banks as possible and that is simply impossible for any one person.

# `bank.IBank` interface

Every bank added to the project _must_ conform to the `bank.IBank` interface such that it can be generically implemented. Additions to the `bank.IBank` interface are expected as edge cases grow however they must be implemented in such a way that the interface remains generic enough to facilitate being adapted to any online banking API.

# Conventions

Banks are added as packages on the top level of this module and use the `{{bank-name}}-{{country-code}}` convention. Country code is using the 3 letter [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3#Current_codes).

## Examples

```
commonwealth-bank-aus
ing-aus
kiwi-bank-nzl
westpac-nzl
```
