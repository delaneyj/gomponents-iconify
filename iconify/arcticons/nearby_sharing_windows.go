package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbySharingWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.975 15.492H6.985c-.8 0-1.485.685-1.485 1.485V40.73c0 .8.685 1.485 1.485 1.485h23.753c.8 0 1.484-.686 1.484-1.485v-4.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.37 20.174v-6.166l9.707 9.707l-9.707 9.706V27.37c-4.91-.114-9.592 2.284-14.16 6.052c1.599-8.336 5.824-13.59 14.16-13.247Zm2.056-9.478a9.699 9.699 0 0 1 9.707 9.707m-9.25-14.618c8.108 0 14.617 6.51 14.617 14.618"/>`),
		g.Group(children),
	)
}