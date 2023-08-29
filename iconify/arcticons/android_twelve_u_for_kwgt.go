package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidTwelveUForKwgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="2.701" height="12.831" x="19.945" y="27.358" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".54"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.786 27.358h-.27a2.161 2.161 0 0 0-2.161 2.162h0v1.434a.27.27 0 0 1-.089.2l-2.443 2.221a.54.54 0 0 0-.036.763q.017.02.036.037l2.443 2.22a.27.27 0 0 1 .089.2v1.435a2.161 2.161 0 0 0 2.16 2.16h.27a.27.27 0 0 0 .27-.27v-4.301a.54.54 0 0 0-.158-.382l-1.271-1.272a.27.27 0 0 1 0-.382l1.271-1.271a.54.54 0 0 0 .159-.382v-4.301a.27.27 0 0 0-.27-.27ZM13.5 24.497h8.206M13.5 10.281l4.103-2.266m0 0v16.482m6.024-11.124a5.488 5.488 0 0 1 5.334-5.563a5.662 5.662 0 0 1 5.539 5.563a5.64 5.64 0 0 1-1.641 3.914c-2.257 1.854-9.232 7.21-9.232 7.21H34.5"/>`),
		g.Group(children),
	)
}