package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AkuhPakistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.05 39.312a11.306 11.306 0 1 1-1.303-19.597"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.949 39.297a9.775 9.775 0 0 1-8.71-9.232a9.875 9.875 0 0 1 7.481-10.318"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.574 19.74a11.572 11.572 0 0 1 6.653 13.035a11.004 11.004 0 0 1-11.498 8.58a11.63 11.63 0 0 1-10.735-9.99"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.617 19.833a9.66 9.66 0 0 1-2.846 12.308a9.903 9.903 0 0 1-12.734-.771"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.992 18.492a11.594 11.594 0 0 1 8.972-11.577a11.069 11.069 0 0 1 12.802 6.55a11.481 11.481 0 0 1-4.429 13.903"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.058 18.412a9.875 9.875 0 0 1 12.436-2.784a9.707 9.707 0 0 1 4.83 11.7"/>`),
		g.Group(children),
	)
}