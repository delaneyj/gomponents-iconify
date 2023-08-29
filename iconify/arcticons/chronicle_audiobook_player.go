package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChronicleAudiobookPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.213 21.06l.011-7.39h13.37v7.389Zm.087 8.522l2.694-2.39l2.592 2.39l2.76-2.39l2.525 2.39l2.827-2.39M15.014 8.221h-2.047a1.76 1.76 0 0 0-1.76 1.76v30.804a1.76 1.76 0 0 0 1.76 1.76h2.047"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.014 8.221v34.324h20.02a1.76 1.76 0 0 0 1.76-1.76V9.982a1.76 1.76 0 0 0-1.76-1.76ZM8.184 18.29h.674v12.228h0h-.674A3.684 3.684 0 0 1 4.5 26.835v-4.86a3.684 3.684 0 0 1 3.684-3.684Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.455H11.667a2.809 2.809 0 0 0-2.809 2.809V18.29m30.958 12.23h-.674h0V18.291h.674a3.684 3.684 0 0 1 3.684 3.683v4.86a3.684 3.684 0 0 1-3.684 3.684ZM24 5.455h12.333a2.809 2.809 0 0 1 2.809 2.809V18.29"/>`),
		g.Group(children),
	)
}