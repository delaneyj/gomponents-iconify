package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminPanelSettingsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17q.625 0 1.063-.438T18.5 15.5q0-.625-.438-1.063T17 14q-.625 0-1.063.438T15.5 15.5q0 .625.438 1.063T17 17Zm0 3q.775 0 1.425-.363t1.05-.962q-.55-.325-1.175-.5T17 18q-.675 0-1.3.175t-1.175.5q.4.6 1.05.963T17 20Zm-5 2q-3.475-.875-5.738-3.988T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375v4.3q-.475-.2-.975-.363T18 10.075V6.4l-6-2.25L6 6.4v4.7q0 1.175.313 2.35t.875 2.238Q7.75 16.75 8.55 17.65t1.775 1.5q.275.8.725 1.525t1.025 1.3q-.025 0-.037.013T12 22Zm5 0q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm-5-10.35Z"/>`),
		g.Group(children),
	)
}