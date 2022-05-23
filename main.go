package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := flag.String("token", "", "discord bot token")
	mode := flag.String("mode", "tvl", "bot type")
	status := flag.Int("status", 0, "0: playing, 1: listening")
	refresh := flag.Int("refresh", 300, "seconds between refresh")
	flag.Parse()

	dg, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if len(guilds) == 0 {
		log.Fatal("Not in any guilds!")
	}

	activity := "alta.finance"
	ticker := time.NewTicker(time.Duration(*refresh) * time.Second)

	for {
		select {
		case <-ticker.C:
			switch *mode {
			case "tvl":
				data, err := GetEarn()
				if err != nil {
					log.Println(err)
					continue
				}

				nickname := fmt.Sprintf("TVL: %.2fM", data.Tvl / 1000000)
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						log.Println(err)
						continue
					} else {
						log.Printf("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}

				activity = "Earn Total Volume"
			case "apr":
				data, err := GetEarn()
				if err != nil {
					log.Println(err)
					continue
				}

				nickname := fmt.Sprintf("APR: %.1f%%", data.MaxAPR * 100)
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						log.Println(err)
						continue
					} else {
						log.Printf("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}

				activity = "Earn Max USDC APR"
			case "treasury":
				data, err := GetTreasury()
				if err != nil {
					log.Println(err)
					continue
				}

				nickname := fmt.Sprintf("Treasury: %.2fM", data.Balance / 1000000)
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						log.Println(err)
						continue
					} else {
						log.Printf("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}

				activity = "Alta Finance Treasury"
			}
			switch *status {
			case 0:
				err = dg.UpdateGameStatus(0, activity)
			case 1:
				err = dg.UpdateListeningStatus(activity)
			}
			if err != nil {
				log.Printf("Unable to set activity: %s\n", err)
			} else {
				log.Printf("Set activity: %s\n", activity)
			}
		}
	}
}
