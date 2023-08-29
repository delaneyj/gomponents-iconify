package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mgen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.991 21.598V28.2a2.201 2.201 0 0 1-2.201 2.201h0a2.194 2.194 0 0 1-1.556-.644"/><rect width="4.402" height="5.833" x="20.589" y="21.598" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.201" transform="rotate(-180 22.79 24.514)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 23.799a2.201 2.201 0 0 1 2.201-2.201h0a2.201 2.201 0 0 1 2.201 2.2v3.633M9.5 21.598v5.833"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.902 23.799a2.201 2.201 0 0 1 2.202-2.201h0a2.201 2.201 0 0 1 2.201 2.2v3.633m13.152-1.111a2.2 2.2 0 0 1-1.913 1.11h0a2.201 2.201 0 0 1-2.2-2.2v-1.431a2.201 2.201 0 0 1 2.2-2.201h0a2.201 2.201 0 0 1 2.202 2.2v.716h-4.403M38.5 27.43V23.8a2.201 2.201 0 0 0-2.201-2.201h0a2.201 2.201 0 0 0-2.202 2.2v3.633m0-3.633v-2.201"/>`),
		g.Group(children),
	)
}