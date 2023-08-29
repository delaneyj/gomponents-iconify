package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M239.32 496.668L97.906 229.358L2.278 155.592L239.32 496.668zm271.421-341.076l-95.628 73.766l-141.415 267.31l237.043-341.076zM391.237 232.22H120.763l135.226 265.54l135.248-265.561v.021zm-289.156-23.244L0 131.28l62.028-72.531l88.113 43.13l-48.038 107.097h-.022zm65.321-99.994l-46.639 104.255h270.474l-46.64-104.255H167.403zm242.473 99.994L512 131.28l-61.985-72.531l-88.156 43.13l47.995 107.097h.021zM344.598 90.042H167.402l-85.81-42.657l77.869-33.145h193.057l77.912 33.145l-85.832 42.657z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}