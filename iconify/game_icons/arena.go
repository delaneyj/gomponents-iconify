package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M76 16C36 36 31 46 31 46c40 160 5 275-15 375c0 0 10 25 90 45V256c0-80-30-240-30-240zm360 0s-30 160-30 240v210c80-20 90-45 90-45c-20-100-55-215-15-375c0 0-5-10-45-30zM226 196c-40 0-90 15-90 15v270s50 15 90 15h60c40 0 90-15 90-15V211s-50-15-90-15h-60z"/>`),
		g.Group(children),
	)
}