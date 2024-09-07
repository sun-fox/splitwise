# Expense Sharing Application

This application allows users to add, split, and manage shared expenses among multiple participants. It helps track balances between users, displaying how much each person owes or is owed by others.

## Features

- Users can add expenses and split them equally, by exact amounts, or by percentages.
- Users can view their balances with others.
- Validation for percentages and exact amounts.
- Support for expense names, notes, and images (optional).
- Option to split by shares, where a user can represent more than one person.
- Ability to simplify transactions to reduce complexity between balances.

## Table of Contents

- [Usage](#usage)
- [Expense Types](#expense-types)
- [Commands](#commands)
- [Optional Features](#optional-features)
- [Sample Input/Output](#sample-inputoutput)

---

## Usage

To run the application, create some users and add expenses for them. Each user should have a unique ID, name, email, and mobile number. You can then input expenses and view balances using the supported commands.

## Expense Types

1. **EQUAL**: The expense is split equally among the participants.
2. **EXACT**: Each participant owes a specific amount.
3. **PERCENT**: Each participant pays a percentage of the total amount.
4. **SHARE** (Optional): Some users may represent more than one person when splitting expenses.

---

## Commands

### 1. Adding an Expense

**Command format**:
- `EXPENSE <user-id-of-person-who-paid> <no-of-users> <space-separated-list-of-users> <EQUAL/EXACT/PERCENT> <optional-share-values>`

- `<user-id-of-person-who-paid>`: ID of the user who paid.
- `<no-of-users>`: The number of users sharing the expense.
- `<space-separated-list-of-users>`: List of user IDs involved in the expense.
- `<EQUAL/EXACT/PERCENT>`: The type of split.
- `<optional-share-values>`: (For EXACT or PERCENT only) Values representing how much each user owes.

**Examples**:
1. Split equally among 4 users:

- `EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL`

Here, each user owes `250` to `User1`.

2. Split by exact amounts:

- `EXPENSE u1 1250 2 u2 u3 EXACT 370 880`

Here, `User2` owes `370` and `User3` owes `880` to `User1`.

3. Split by percentages:

- `EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20`

Here, `User1` owes `480`, `User2` owes `240`, and `User3` owes `240` to `User4`.

### 2. Show Balances for All Users

**Command format**:

Displays all balances between users with non-zero amounts.

**Example output**:
User2 owes User1: 620 
User3 owes User1: 1130
User1 owes User4: 230 
User2 owes User4: 240 
User3 owes User4: 240

### 3. Show Balance for a Single User

**Command format**:
`SHOW <user-id>`


Displays balances for the specific user, showing what they owe or are owed.

**Example output**:
User1 owes User4: 230 
User2 owes User1: 620 
User3 owes User1: 1130

---

## Optional Features

1. **Expense Name, Notes, and Images**:
   - Add a descriptive name or notes while adding expenses.
   
2. **Split by Share**:
   - Instead of using percentages, split by shares where each user represents more than one person.
   - **Command format**:
     ```
     EXPENSE <user-id> <amount> <no-of-users> <space-separated-list-of-users> SHARE <share-values>
     ```
   - **Example**:
     ```
     EXPENSE u4 1200 4 u1 u2 u3 u4 SHARE 2 1 1 1
     ```

3. **Simplify Balances**:
   - When simplify is turned on, the app reduces the number of transactions between users.
   - Example: If `User1` owes `250` to `User2` and `User2` owes `200` to `User3`, the system simplifies it to `User1` owes `50` to `User2` and `200` to `User3`.

4. **Show Passbook**:
   - Display all past transactions a user was part of.

---

## Sample Input/Output

### Input
`EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL EXPENSE u1 1250 2 u2 u3 EXACT 370 880 EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20 SHOW SHOW u1`

### Output
For `SHOW`:
User2 owes User1: 620 
User3 owes User1: 1130 
User1 owes User4: 230 
User2 owes User4: 240 
User3 owes User4: 240

For `SHOW u1`:
User2 owes User1: 620 
User3 owes User1: 1130 
User1 owes User4: 230
---

## Validations

1. **Percentage Validation**:
   - Ensure that the sum of percentages in a `PERCENT` expense equals 100.
   
2. **Exact Amount Validation**:
   - Ensure that the sum of the amounts in an `EXACT` expense equals the total amount.

3. **Rounding**:
   - Round off amounts to two decimal places.

---

## Future Enhancements

- Add user authentication.
- Enable filtering of passbook by date range.
- Include notification system for dues reminders.

---

Feel free to build and expand this application as per your requirements!

--- 