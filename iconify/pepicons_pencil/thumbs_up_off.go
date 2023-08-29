package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.7 8.503H3.2v5.5h2.5v-5.5Zm-2.5-1a1 1 0 0 0-1 1v5.5a1 1 0 0 0 1 1h2.5a1 1 0 0 0 1-1v-5.5a1 1 0 0 0-1-1H3.2ZM9.744 2.83c.3.05.501.333.451.632l-.223 1.342a5.12 5.12 0 0 1-2.21 3.418l-.61-.914a4.022 4.022 0 0 0 1.736-2.685l.224-1.342a.55.55 0 0 1 .632-.451Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.416 3.164a.733.733 0 0 0-1.23.34l-1.065-.266c.345-1.38 2.065-1.857 3.071-.85l.047.046a.55.55 0 1 1-.777.777l-.046-.047Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.992 6.65a.55.55 0 0 1-.4-.665l.163-.653c.062-.246.089-.5.08-.753l-.015-.451a1.481 1.481 0 0 0-.378-.939l.817-.733c.405.45.638 1.029.659 1.634l.015.451a3.82 3.82 0 0 1-.112 1.058l-.163.652a.55.55 0 0 1-.666.4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.85 6.118a.55.55 0 0 1 .55-.55h1.647a.55.55 0 0 1 0 1.099H12.4a.55.55 0 0 1-.55-.55Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.94 7.837a2.117 2.117 0 0 0-1.893-1.17V5.569c1.218 0 2.332.688 2.876 1.777l.362.723a.55.55 0 0 1-.983.491l-.361-.723Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.05 14.32a.55.55 0 0 1-.32-.707l.594-1.568a4.832 4.832 0 0 0-.04-3.524l1.019-.412a5.93 5.93 0 0 1 .049 4.325l-.594 1.568a.55.55 0 0 1-.708.319Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.098 14.926a2.122 2.122 0 0 0 2.385-.93l.293-.477a.55.55 0 1 1 .936.576l-.293.477a3.22 3.22 0 0 1-3.62 1.411l.3-1.057Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.918 14.248a.55.55 0 0 1 .646-.43l5.493 1.098a.55.55 0 0 1-.216 1.077L7.35 14.895a.55.55 0 0 1-.43-.646Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.46 7.216a.55.55 0 0 1 .55.55v6.59a.55.55 0 1 1-1.098 0v-6.59a.55.55 0 0 1 .549-.55Z" clip-rule="evenodd"/><path d="M5.45 12.504a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}