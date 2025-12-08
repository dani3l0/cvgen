package config

const defaultPromptAboutMe = `
- Full Name: Jonathan Wick
- Phone: +48 347 374 743
- Email: no@rep.ly
- Location: Warsaw, PL

I am a true Python programmer. I have more than 50 projects on GitHub, invented AI CV generation engine and been on scholarship at NASA in Poland.
I have 1 year of commercial experience in Django programming. Also, I'm familiar with SQL databases.
`

const defaultPromptHtmlResponseFormat = `
HTML resume demands:
- add CSS so it looks better
- a two-column style
- do not exceed A4 dimensions
- add 'margin: 0; padding: 0' to <body>
- light-red pastel background on the left column
- try to not write very long texts on one side - it will exceed A4 dimensions
- use emojis at short personal information like contact, location
`

const defaultPromptNotes = `
Generate HTML contents strictly respecting above prompt. Once finished, check again and correct possible caveats.
`
