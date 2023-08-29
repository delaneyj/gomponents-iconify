package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cubit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m212.815 340.95l-65.444-39.464v-78.309l65.444 37.885v79.887zM84.17 171.375l135.46-81.681l135.234 81.546l75.891-43.932L219.628 0L8.278 127.444l75.89 43.932zm202.092 39.577l-66.633-40.18l-66.858 40.315l66.629 38.57l66.862-38.705zm76.184-28.886v157.95l-136.462 82.287V512L440 383.395V137.603l-77.555 44.464zm-149.63 239.96l-136.004-82.01V182.332L0 137.868v245.527l212.815 128.328v-89.695zm79.07-199.115l-65.902 38.15v80.164l65.902-39.74v-78.574z"/>`),
		g.Group(children),
	)
}