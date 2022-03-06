package data

var DefaultStarterFile = `---
# user's hierarchy groups to manage level of access for the protected files
# the first group in the group list will be the DEFAULT GROUP
groups:
  dev:
    - .env.dev
    - folder/secret.txt
  staging:
    - .env.stage
    - folder/another-secret.txt

  # if you want a group to have all of the files that another group have just assign it with a 'group.' prefix. 
  # example: '- group.anotherGroupName'.
  prod:
    - group.staging
    - .env.prod

# list of users and its incorporated group
users:
  dev:
    - 'user1@gmail.com'
  staging:
    - 'user1@gmail.com'
  prod:
    - 'user1@gmail.com'
`
