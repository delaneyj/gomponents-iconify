package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlitzortungLightningMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.622 33.893c5.642 0 10.216 1.59 10.216 3.55s-4.574 3.548-10.216 3.548s-10.216-1.589-10.216-3.549c0-.98 1.144-1.867 2.992-2.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.034 31.671c9.077 0 16.436 2.648 16.436 5.915S33.11 43.5 24.033 43.5S7.572 40.852 7.6 37.586c.024-2.853 4.155-4.985 8.04-5.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.482 35.972c2.396 0 3.513.666 3.513 1.398s-1.978 1.613-4.373 1.613s-4.288-.88-4.301-1.613c-.013-.684 1.538-1.046 2.059-1.188m-5.178-24.226c-5.902-7.316-12.46 5.286-6.094 6.237c8.78 1.312 18.331 1.743 26.31.359c8.017-1.391 2.838-12.632-3.655-6.31c4.441-2.92-3.15-13.111-7.85-2.15c3.412-6.81-10.098-8.723-8.71 1.864Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.153 19.11l-3.385 4.567l.359 2.725l-2.115 1.72"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.768 23.677c-2.178-.283-4.246-.257-5.95.79m7.628-3.053l5.706 4.952l-1.756 5.735l-2.258 2.258m4.014-7.993l3.692 1.111l3.011-2.51l1.9.144m-4.91 2.366l-.072 2.187l1.075.86m-6.452 1.577l3.62 5.592"/>`),
		g.Group(children),
	)
}