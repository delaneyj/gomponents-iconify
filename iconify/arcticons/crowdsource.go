package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crowdsource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.98" cy="10.817" r="2.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.98" cy="22.674" r="2.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.491" cy="16.849" r="2.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.469" cy="16.849" r="2.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.3 39.523v-7.8c-.209-1.04-.313-2.705-1.145-4.057c-2.392-3.744-7.488-7.384-10.297-10.4c-1.56-1.769-4.472.935-2.912 2.6l4.472 6.032c3.329 4.68-1.352 9.57 9.881 13.626Zm9.36 0v-7.8c.208-1.04.312-2.705 1.144-4.057c2.393-3.744 7.49-7.384 10.297-10.4c1.56-1.769 4.577.935 2.913 2.704l-4.473 6.032c-3.328 4.577 1.352 9.465-9.88 13.522Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.3 31.93c-.105-4.368 9.256-4.576 9.36-.207m-9.36 7.8c2.912.625 6.032.625 9.36 0"/>`),
		g.Group(children),
	)
}