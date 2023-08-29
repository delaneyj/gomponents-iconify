package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsFootballOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12Zm-2.95 6.95l-4-4q-.075.95-.037 1.975T5.2 18.75q.575.175 1.738.225t2.112-.025Zm2.4-.4q1.475-.325 2.65-.925t2.05-1.475q.85-.85 1.45-2.013t.95-2.637L12.5 5.45q-1.425.35-2.575.963T7.9 7.9q-.875.875-1.488 2.038T5.45 12.55l6 6ZM9.9 15.5l-1.4-1.4l5.6-5.6l1.4 1.4l-5.6 5.6Zm9.05-6.4q.1-.975.063-2.025T18.8 5.25q-.575-.2-1.737-.25t-2.113.05l4 4.05ZM7.75 21q-1.425 0-2.6-.213T3.7 20.3q-.275-.3-.488-1.5T3 16.15q0-2.975.9-5.512T6.45 6.45Q8.1 4.8 10.675 3.9T16.25 3q1.45 0 2.613.213T20.3 3.7q.275.3.488 1.5T21 7.9q0 2.925-.9 5.463t-2.55 4.187q-1.625 1.625-4.2 2.538T7.75 21Z"/>`),
		g.Group(children),
	)
}