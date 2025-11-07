{{/* TRIGGER TYPE: COMMAND */}}
{{/* TRIGGER: `-curse` */}}
{{/* CHANNELS: ALL */}}
{{/* ROLES: ALL */}}

{{$args := parseArgs 1 "" ( carg "string" "command_name" )}}
{{$command := $args.Get 0}}

{{$role_spicy := "SPICY"}}
{{$role_gag_curse := "Gag Curse"}}
{{$role_consent := "Consent to Mute"}}

{{$has_spicy := hasRoleName $role_spicy}}
{{$has_curse := hasRoleName $role_gag_curse}}

{{if eq $has_spicy false}}
	{{if eq $has_curse true}}
		{{removeRoleName $role_gag_curse}}
		{{$has_curse = false}}
	{{end}}
{{end}}

{{/* Set the authentic sticker for the "you don't have the NSFW role" message */}}
{{/* If you don't have one, you can use 754109772449710080, that's the "Clyde Bot - No" (a robot shaking his head and pointing a finger at you) */}}
{{$sticker_id := 1375885651676893326}}

{{$msg_no_spicy := ( printf "You don't have the 🌶️ **NFWS role** (**@%s**) to have the Gag Curse!" $role_spicy )}}
{{$msg_curse_lifted := ( printf "**The Gag Curse** was **lifted** from %s! 😜" .User.Mention )}}
{{$msg_curse_applied := ( printf "**The Gag Curse** was **applied** on %s! 😷" .User.Mention )}}

{{if eq $command "status"}}
	{{$not := "not "}}
	{{$emoji := "😜"}}
  	{{if eq $has_curse true}}
		{{$not = ""}}
		{{$emoji = "😷"}}
	{{end}}
	{{sendMessage nil ( printf "%s's **Gag Curse** status: **%scursed**! %s" .User.Mention $not $emoji )}}
{{else if eq $command "toggle"}}
	{{if eq $has_spicy false}}
		{{sendMessage nil ( complexMessage "sticker" $sticker_id )}}
		{{sendMessage nil ( $msg_no_spicy )}}
		{{return}}
	{{end}}
	
	{{if eq $has_curse true}}
		{{removeRoleName $role_gag_curse}}
		{{sendMessage nil ( $msg_curse_lifted )}}
	{{else}}
		{{addRoleName $role_gag_curse}}
		{{sendMessage nil ( $msg_curse_applied )}}
	{{end}}
{{else if eq $command "off"}}
	{{if eq $has_curse true}}
		{{removeRoleName $role_gag_curse}}
		{{sendMessage nil ( $msg_curse_lifted )}}
	{{else}}
		{{sendMessage nil ( printf "There's **no curse** to begin with, %s! 🤔" .User.Mention )}}
	{{end}}
{{else if eq $command "on"}}
	{{if eq $has_spicy false}}
		{{sendMessage nil ( complexMessage "sticker" $sticker_id )}}
		{{sendMessage nil ( $msg_no_spicy )}}
		{{return}}
	{{end}}
	
	{{if eq $has_curse false}}
		{{addRoleName $role_gag_curse}}
		{{sendMessage nil ( $msg_curse_applied )}}
	{{else}}
		{{sendMessage nil ( printf "You're already **cursed**, %s! 😷" .User.Mention )}}
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
* Only users with the 🌶️ **NFWS role** (**@{{$role_spicy}}**) can have this "curse".
* When gagged, you're encouraged to modify your speech accordingly for a short while (gag-talk, gag-speech, mmph-ing, you get it).
* Only if you're **@{{$role_consent}}**, you'll receive a __1 minute timeout__ upon gagging (see Reaction Roles channel).
{{else}}
	{{sendMessage nil ( "Invalid command! Available commands are: `status` `toggle` `on` `off` `help`" )}}
{{end}}
