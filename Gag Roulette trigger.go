{{/* TRIGGER TYPE: REGEX */}}
{{/* TRIGGER: `.*` */}}
{{/* CHANNELS: `#😷the-gag-curse` */}}
{{/* ROLES: ALL */}}

{{$is_consent := hasRoleName "Consent to Mute"}}
{{$is_admin := hasPermissions .Permissions.ModerateMembers}}
{{$can_mute := true}}
{{if eq $is_consent false}}
	{{$can_mute = false}}
{{else if eq $is_admin true}}
	{{$can_mute = false}}
{{end}}

{{/* Set the actual ID of the gagging script on your server instead */}}
{{$gagging_cmd_id := 9}}
{{execCC $gagging_cmd_id nil 0 (sdict "GaggingType" "Roulette" "UserID" .User.ID "UserMention" .User.Mention "CanMute" $can_mute)}}
