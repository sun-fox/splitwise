Expense Sharing Application
Overview
This application helps users manage and split expenses among a group of people. Each user can add expenses, select how to split them (equally, exactly, or by percentage), and the app will keep track of who owes whom and by how much.

Features
User Management: Create and manage user profiles.
Expense Management: Add expenses and split them among users.
Balance Tracking: View balances and transactions for users.
Flexible Splitting Options: Split expenses equally, exactly, or by percentage.
Requirements
User
Each user should have the following attributes:

userId: Unique identifier
name: Full name
email: Email address
mobile number: Contact number
Expense
Expenses can be split in three ways:

EQUAL: Split the expense equally among all selected users.
EXACT: Specify exact amounts for each user.
PERCENT: Specify percentage shares for each user (should total 100%).
Validation
Percent Split: Verify that the total percentage is 100.
Exact Split: Verify that the total amount matches the specified amounts.
Rounding
Amounts should be rounded to two decimal places. For example, if an amount is split equally among three users, one user might get slightly more than the others to ensure the total amount is accurate.

Inputs
Creating Users
You can create users in the main method. Example:

bash
Copy code
User1 (id: u1)
User2 (id: u2)
User3 (id: u3)
User4 (id: u4)
Adding Expenses
Expenses are recorded using the following format:

php
Copy code
EXPENSE <user-id-of-person-who-paid> <no-of-users> <space-separated-list-of-users> <EQUAL/EXACT/PERCENT> <space-separated-values-in-case-of-non-equal>
Viewing Balances
For All Users:

sql
Copy code
SHOW
For a Single User:

sql
Copy code
SHOW <user-id>
Output
For Single User
When showing balances for a single user, output balances where the user is involved:

php
Copy code
<user-id-of-x> owes <user-id-of-y>: <amount>
If there are no balances, print:

yaml
Copy code
No balances
If the user owes money, they’ll be listed as <user-id-of-x>. If they are owed money, they’ll be <user-id-of-y>.

Optional Features
Expense Name and Notes: Add names, notes, or images to expenses.
Share Splitting: Split expenses by specified shares (e.g., SHARE 2 1 1 1).
Passbook: View a list of all transactions a user was part of.
Simplify Expenses: Simplify balances to reduce the number of transactions. For example, if User1 owes User2, and User2 owes User3, it could be simplified to User1 owing User3 directly.
Examples
Example 1: Equal Split
Input:

yaml
Copy code
u1 1000 4 u1 u2 u3 u4 EQUAL
Output:

User2 owes User1: 250
User3 owes User1: 250
User4 owes User1: 250
Example 2: Exact Split
Input:

yaml
Copy code
u1 1250 2 u2 u3 EXACT 370 880
Output:

User2 owes User1: 370
User3 owes User1: 880
Example 3: Percent Split
Input:

sql
Copy code
u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20
Output:

User1 owes User4: 480
User2 owes User4: 240
User3 owes User4: 240