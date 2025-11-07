{{/* TRIGGER TYPE: COMMAND */}}
{{/* TRIGGER: `-curse` */}}
{{/* CHANNELS: ALL */}}
{{/* ROLES: ALL */}}

{{$args := parseArgs 1 "" ( carg "string" "command_name" )}}
{{$command := $args.Get 0}}

{{$has_spicy := hasRoleName "SPICY"}}
{{$has_curse := hasRoleName "Gag Curse"}}

{{if eq $has_spicy false}}
	{{if eq $has_curse true}}
		{{removeRoleName "Gag Curse"}}
		{{$has_curse = false}}
	{{end}}
{{end}}

{{if eq $command "status"}}
	{{$not := "not "}}
	{{$emoji := "😜"}}
  	{{if eq $has_curse true}}
		{{$not = ""}}
		{{$emoji = "😷"}}
	{{end}}
{{.User.Mention}}'s **Gag Curse** status: **{{$not}}cursed**! {{$emoji}}
{{else if eq $command "toggle"}}
	{{if eq $has_spicy false}}
		{{sendMessage nil ( complexMessage "sticker" 1375885651676893326 )}}
You don't have the 🌶️ **NFWS role** (**@SPICY**) to have the Gag Curse!
		{{return}}
	{{end}}
	
	{{if eq $has_curse true}}
		{{removeRoleName "Gag Curse"}}
**The Gag Curse** was **lifted** from {{.User.Mention}}! 😜
	{{else}}
		{{addRoleName "Gag Curse"}}
**The Gag Curse** was **applied** on {{.User.Mention}}! 😷
	{{end}}
{{else if eq $command "off"}}
	{{if eq $has_curse true}}
		{{removeRoleName "Gag Curse"}}
**The Gag Curse** was **lifted** from {{.User.Mention}}! 😜
	{{else}}
There's **no curse** to begin with, {{.User.Mention}}! 🤔
	{{end}}
{{else if eq $command "on"}}
	{{if eq $has_spicy false}}
		{{sendMessage nil ( complexMessage "sticker" 1375885651676893326 )}}
You don't have the 🌶️ **NFWS role** (**@SPICY**) to have the Gag Curse!
		{{return}}
	{{end}}
	
	{{if eq $has_curse false}}
		{{addRoleName "Gag Curse"}}
**The Gag Curse** was **applied** on {{.User.Mention}}! 😷
	{{else}}
You're already **cursed**, {{.User.Mention}}! 😷
	{{end}}
{{else if eq $command "help"}}
# The Gag Curse 😷
All those who have the curse active have a small chance to be **"gagged"** for a short while after sending any message!

Command format: `-curse subcommand`, where subcommand is one of the available commands.
Available commands are: 
* `status` : get your curse status.
* `toggle` : toggle your curse from on to off and vice versa.
* `on` : activate your curse.
* `off` : deactivate your curse.
* `help` : information about the Gag Curse.

**Notes:**
* Only users with the 🌶️ **NFWS role** (**@SPICY**) can have this "curse".
* When gagged, you're encouraged to modify your speech accordingly for a short while (gag-talk, gag-speech, mmph-ing, you get it).
* Only if you're **@Consent to Mute**, you'll receive a __1 minute timeout__ upon gagging (see Reaction Roles channel).
{{else}}
Invalid command! Available commands are: `status` `toggle` `on` `off` `help`
{{end}}
