package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SprayingNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSprayingNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM10 20a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h.054c.257 2.35 1.42 4.28 3.156 5.65c1.957 1.543 4.602 2.35 7.45 2.35c4.552 0 8.76-1.833 10.172-6.007a2.001 2.001 0 0 0 1.845-3.084l3.788-2.062l-.956-1.756l-4.856 2.643L27 30.364l-1.351-.507h-.002A1 1 0 0 1 25 28.92V26a2 2 0 1 0-4 0v5a2 2 0 0 0 1.298 1.873l3.913 1.467l2.723 1.021C27.897 38.388 24.771 40 20.66 40c-2.487 0-4.671-.707-6.21-1.921c-1.313-1.035-2.184-2.456-2.398-4.218A4.002 4.002 0 0 0 15 30v-6a4 4 0 0 0-4-4h1.5a1 1 0 1 0 0-2h-4a1 1 0 1 0 0 2H10Zm16.465-10A3.998 3.998 0 0 0 23 8c-1.48 0-2.773.804-3.465 2h6.93ZM25 15v-3h3v4h-2a1 1 0 0 1-1-1Zm5 2v-5h2v-2h-3.341a6 6 0 1 0-3.892 7.736c.376.17.794.264 1.233.264h4v-1Zm10.152 7.47l-2.5 4l-.956 1.53H42v-2h-1.696l1.544-2.47l-1.696-1.06ZM18 20a2 2 0 0 0-2 2v9.5a5.5 5.5 0 0 0 8.758 4.431l-3.162-1.186A4 4 0 0 1 19 31v-5a4 4 0 0 1 8 0v-3a3 3 0 0 0-3-3h-6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSprayingNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}