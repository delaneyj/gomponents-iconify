package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsWheelchairNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm25.397 26.095a7.98 7.98 0 0 1 2.602 6.036A8 8 0 0 1 12 32a8 8 0 0 1 13.397-5.905ZM30 31.988l.62-.045s3.054 5.365 4.285 7.224c1.231 1.86 4.333.31 3.434-1.764c-.9-2.074-4.239-9.942-4.239-9.942s-.26-.787-1.397-.787c-.6 0-2.647-.165-4.45-.322a10.054 10.054 0 0 0-2.845-2.765l.145-1.92s1.468 2.409 2.344 2.91c.876.5 7.363 1.144 8.547 1.191c1.184.048 2.344-2.479.497-3.123c-1.846-.643-6.677-1.788-6.677-1.788s-3.835-7.581-6.108-7.796c-2.273-.214-2.676 0-4.736 2.909c-.363.512-.668 1.193-.92 1.977V15.75l-3.4-2.55l-1.2 1.6l2.6 1.95v5.88A10.004 10.004 0 0 0 10 32c0 5.523 4.477 10 10 10s10-4.477 10-10v-.012Zm-5.37.385l1.364-.098a6 6 0 0 0-8.476-5.739c.043 1.15.162 2.076.34 2.57c1.363 3.791 6.771 3.267 6.771 3.267ZM30.5 9.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsWheelchairNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}