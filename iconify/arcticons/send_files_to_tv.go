package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendFilesToTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.5 24h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.945 11.913l4.756-2.277l2.278 4.756m-2.278-4.756L17.639 24m12.416 12.087l-4.756 2.277l-2.278-4.756m2.278 4.756L30.361 24"/>`),
		g.Group(children),
	)
}