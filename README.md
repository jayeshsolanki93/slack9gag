# slack9gag

Browse 9gag right from Slack (Uses [infinigag](https://github.com/k3min/infinigag))

![](https://imgur.com/u9ycoHY.gif)

## Usage

From any Slack channel, just type `/9gag [section] [sub-section]`. The content will be visible just to you.

#### Sections:
 - cute
 - comic
 - cosplay
 - design
 - food
 - funny
 - geeky
 - gif
 - girl
 - meme
 - nsfw
 - timely
 - wtf

#### Sub-sections:
 - fresh
 - hot

## Integrate with your team

1. Go to your channel
2. Click on **Apps and integration** and find **Slash Commands**.
3. Click **Install** on front of a team name to install slash command.
  - Command: `/9gag`
  - URL: `slack9gag.herokuapp.com`
  - Method: `POST`
  - For the **Autocomplete help text**, check to show the command in autocomplete list.
    - Description: `Browse 9gag right from Slack.`
    - Usage hint: `[section] [subsection]`
  - Descriptive Label: `Browse 9gag`

## Contributing

- Please use the [issue tracker](https://github.com/jayeshsolanki93/slack9gag/issues) to report any bugs or file feature requests.
