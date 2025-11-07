{{/* TRIGGER TYPE: NONE */}}
{{/* TRIGGER: Can only be triggered manually by other custom commands */}}

{{/* Triggers are sending a dictionary with neccessary data */}}
{{$gagging_type := .ExecData.GaggingType}}
{{$user_id := .ExecData.UserID}}
{{$user_mention := .ExecData.UserMention}}
{{$can_mute := .ExecData.CanMute}}

{{$chance := 0}}
{{/* Set the chances below to be gagged in X.X% (35 = 3.5%) */}}
{{if eq $gagging_type "Roulette"}}
	{{$chance = 750}}
{{else if eq $gagging_type "Curse"}}
	{{$chance = 25}}
{{end}}

{{/* Calculating the probability of receiving a gag in the roulette in % */}}
{{$calced_prob_receive_gag := mult (fdiv $chance 1000) 100}}

{{$roll := randInt 1000}}
{{if lt $roll $chance}}
	{{/* Defining gag categories */}}
	{{$gags_categories := sdict "exotic" 10 "deliberate" 40 "improvised" 50}}
	
	{{/* Calculating total weights for gags categories */}}
	{{$weights := 0}}
	{{range $gag_cat, $prob_weight := $gags_categories}}
    	{{$weights = add $weights $prob_weight}}
	{{end}}
	
	{{/* Selecting the gag category */}}
	{{$gag_category := ""}}
	{{$calced_prob_category := 1.0}}
	{{$roll = randInt $weights}}
	{{$chance = 0}}
	{{range $gag_cat, $prob_weight := $gags_categories}}
		{{$chance = add $chance $prob_weight}}
     	{{if lt $roll $chance}}
     		{{$gag_category = $gag_cat}}
			{{/* Calculating the selected gag category probability */}}
			{{$calced_prob_category = fdiv $prob_weight $weights}}
     		{{break}}
     	{{end}}
	{{end}}
	
	{{/* Defining the list of gags according to the selected gag category */}}
	{{$gags := sdict "gag_kind" 1}}
	{{$a := "an"}}
	{{if eq $gag_category "exotic"}}
		{{$gags = sdict "glue" 2 "slime" 1 "honey" 1 "apple" 1 "carrot" 1 "lemon" 1 "cucumber" 1 "stick grenade" 1}}
	{{else if eq $gag_category "improvised"}}
		{{$gags = sdict "tape wrap" 3 "tape strip" 4 "microfoam wrap" 1 "microfoam strip" 2 "OTM cloth" 4 "cloth cleave" 4 "cloth knot" 3}}
	{{else if eq $gag_category "deliberate"}}
		{{$gags = sdict "ball" 4 "wiffle" 4 "bit" 2 "panel" 3 "plug" 2 "ring" 3 "inflatable" 1 "penis" 1 "gwen hood" 1}}
		{{$a = "a"}}
	{{end}}
	
	{{/* Calculating total weights for gags list */}}
	{{$weights = 0}}
	{{range $gag_type, $prob_weight := $gags}}
    	{{$weights = add $weights $prob_weight}}
	{{end}}
	
	{{/* Selecting the gag kind */}}
	{{$gag := ""}}
	{{$calced_prob_kind := 0}}
	{{$roll = randInt $weights}}
    {{$chance = 0}}
	{{range $gag_type, $prob_weight := $gags}}
		{{$chance = add $chance $prob_weight}}
     	{{if lt $roll $chance}}
     		{{$gag = $gag_type}}
			{{/* Calculating the selected gag kind probability */}}
			{{$calced_prob_kind = fdiv $prob_weight $weights}}
     		{{break}}
     	{{end}}
	{{end}}
	
	{{/* Defining the list of mouth stuffings for an improvised gag */}}
	{{$stuffings := sdict "no" 30 "a rag" 25 "a hot sauce coated rag" 3 "a sponge" 20 "a sock" 20 "a hot sauce coated sock" 2 "a panty" 15 "a hot sauce coated panty" 2 "a stress ball" 10}}
	
	{{/* Calculating total weights for stuffings */}}
	{{$weights = 0}}
	{{range $stuffing_kind, $prob_weight := $stuffings}}
    	{{$weights = add $weights $prob_weight}}
	{{end}}
	
	{{/* Selecting the stuffing of an improvised gag and calculating its probability */}}
	{{$stuffing := ""}}
	{{$calced_prob_stuffing := 1.0}}
	{{$roll = randInt $weights}}
    {{$chance = 0}}
	{{if eq $gag_category "improvised"}}
		{{range $stuffing_kind, $prob_weight := $stuffings}}
			{{$chance = add $chance $prob_weight}}
     		{{if lt $roll $chance}}
     			{{$stuffing = $stuffing_kind}}
				{{/* Calculating the stuffing probability */}}
				{{$calced_prob_stuffing = fdiv $prob_weight $weights}}
     			{{break}}
     		{{end}}
		{{end}}
	{{end}}
	
	{{/* Calculating the total chance of the selected type of gag being chosen in % */}}
	{{$calced_prob_gag := mult $calced_prob_category $calced_prob_kind $calced_prob_stuffing 100.0}}
	
	{{/* The chance of being gagged and the chance of choosing a certain kind of gag are separated on purpose */}}
	{{$gag_msg_header := ""}}
	{{if eq $gagging_type "Roulette"}}
		{{$gag_msg_header = printf "**%s wins a gag!** [Win chance: %.2f%%]" $user_mention $calced_prob_receive_gag}}
	{{else if eq $gagging_type "Curse"}}
		{{$gag_msg_header = printf "**The Gag Curse has been activated on %s!**" $user_mention}}
	{{end}}
	
	{{$title := ""}}
	{{if eq $gagging_type "Roulette"}}
		{{$title = "winner"}}
	{{else if eq $gagging_type "Curse"}}
		{{$title = "cursed one"}}
	{{end}}
	{{$stuffing_str := ""}}
	{{if eq $gag_category "improvised"}}
		{{$stuffing_str = printf " with **%s** mouth stuffing" $stuffing}}
	{{end}}
	{{$gag_msg_main := printf "The %s has been gagged with %s **%s** type of gag, the **%s** gag%s! [Item chance: %.2f%%]" $title $a $gag_category $gag $stuffing_str $calced_prob_gag}}
	
	{{$gag_msg_cantgag := ""}}
	{{if eq $can_mute true}}
		{{$silent := execAdmin "Timeout" $user_id "1m" "The Gag Curse!"}}
	{{else}}
		{{$gag_msg_cantgag = "-# (Can't actually gag you. Pretend to be gagged for a minute, would you? Mmph hmph?)"}}
		{{$roll = randInt 100}}
		{{if lt $roll 10}}
			{{$gag_msg_cantgag = "-# (**Gag-talk matafaka, do you speak it?!** 🔫)"}}
		{{end}}
	{{end}}
	
    {{/* Set the authentic sticker for the gagging message */}}
    {{/* If you don't have one, you can use 773904449440579624, that's the "Lonely Leif - Eating" (a blushing leaf with a mouthful of ice cream ) */}}
    {{$sticker_id := 1375929269217136690}}
	{{sendMessage nil ( complexMessage "sticker" $sticker_id )}}
	{{sendMessage nil ( joinStr "\n" (cslice $gag_msg_header $gag_msg_main $gag_msg_cantgag) )}}
{{else}}
	{{if eq $gagging_type "Roulette"}}
		{{/* Calculating the probability of NOT receiving a gag in the roulette in % */}}
		{{$calced_prob_no_gag := sub 100.0 $calced_prob_receive_gag}}
		
		{{$no_gag_msg := printf "No gag, %s? 🤨 [No gag chance: %.2f%%]" $user_mention $calced_prob_no_gag}}
		{{sendMessage nil $no_gag_msg}}
	{{end}}
{{end}}
