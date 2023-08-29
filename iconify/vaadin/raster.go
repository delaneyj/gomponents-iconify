package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 7h1v1H7V7zM5 7h1v1H5V7zM3 7h1v1H3V7zM1 7h1v1H1V7zm5-1h1v1H6V6zM4 6h1v1H4V6zM2 6h1v1H2V6zM0 6h1v1H0V6zm7-1h1v1H7V5zM5 5h1v1H5V5zM3 5h1v1H3V5zM1 5h1v1H1V5zm5-1h1v1H6V4zM4 4h1v1H4V4zM2 4h1v1H2V4zM0 4h1v1H0V4zm7-1h1v1H7V3zM5 3h1v1H5V3zM3 3h1v1H3V3zM1 3h1v1H1V3zm5-1h1v1H6V2zM4 2h1v1H4V2zM2 2h1v1H2V2zM0 2h1v1H0V2zm7-1h1v1H7V1zM5 1h1v1H5V1zM3 1h1v1H3V1zM1 1h1v1H1V1zm5-1h1v1H6V0zM4 0h1v1H4V0zM2 0h1v1H2V0zM0 0h1v1H0V0zm15 7h1v1h-1V7zm-2 0h1v1h-1V7zm-2 0h1v1h-1V7zM9 7h1v1H9V7zm5-1h1v1h-1V6zm-2 0h1v1h-1V6zm-2 0h1v1h-1V6zM8 6h1v1H8V6zm7-1h1v1h-1V5zm-2 0h1v1h-1V5zm-2 0h1v1h-1V5zM9 5h1v1H9V5zm5-1h1v1h-1V4zm-2 0h1v1h-1V4zm-2 0h1v1h-1V4zM8 4h1v1H8V4zm7-1h1v1h-1V3zm-2 0h1v1h-1V3zm-2 0h1v1h-1V3zM9 3h1v1H9V3zm5-1h1v1h-1V2zm-2 0h1v1h-1V2zm-2 0h1v1h-1V2zM8 2h1v1H8V2zm7-1h1v1h-1V1zm-2 0h1v1h-1V1zm-2 0h1v1h-1V1zM9 1h1v1H9V1zm5-1h1v1h-1V0zm-2 0h1v1h-1V0zm-2 0h1v1h-1V0zM8 0h1v1H8V0zM7 15h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7V9zM5 9h1v1H5V9zM3 9h1v1H3V9zM1 9h1v1H1V9zm5-1h1v1H6V8zM4 8h1v1H4V8zM2 8h1v1H2V8zM0 8h1v1H0V8zm15 7h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1V9zm-2 0h1v1h-1V9zm-2 0h1v1h-1V9zM9 9h1v1H9V9zm5-1h1v1h-1V8zm-2 0h1v1h-1V8zm-2 0h1v1h-1V8zM8 8h1v1H8V8z"/>`),
		g.Group(children),
	)
}