package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336zm-896 779q39 0 77-55l222-312v671q0 31 21.5 52.5t52.5 21.5h149q31 0 53-21.5t22-52.5V373q0-31-22-52.5t-53-21.5h-149q-13 0-24.5 2t-20 4.5t-17.5 8.5t-14 9t-12.5 11t-10 11t-9.5 12t-8 10L896 729L639 367q-1-1-8-10t-9.5-12t-10-11t-12.5-11t-14-9t-17.5-8.5t-20-4.5t-24.5-2H374q-31 0-53 21.5T299 373v1046q0 31 22 52.5t53 21.5h149q31 0 52.5-21.5T597 1419V748l222 312q38 55 77 55z"/>`),
		g.Group(children),
	)
}