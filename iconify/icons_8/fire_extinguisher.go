package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4c-1.292 0-2.394.844-2.813 2H12c-2.745 0-5 2.255-5 5h2c0-1.655 1.345-3 3-3h1v2.47c-.32.237-.733.575-1.22 1.06C10.954 12.36 10 13.5 10 15v13h12V15c0-1.5-.953-2.64-1.78-3.47c-.487-.485-.9-.823-1.22-1.06v-.283l4.844.813l1.156.188V4.812L23.844 5l-5.094.844C18.292 4.77 17.234 4 16 4zm0 2c.555 0 1 .445 1 1v3h-2V7c0-.555.445-1 1-1zm7 1.188v1.625l-4-.688v-.25l4-.688zM14.375 12h3.25c.15.105.578.39 1.156.97c.673.67 1.22 1.53 1.22 2.03v11h-8V15c0-.5.547-1.36 1.22-2.03c.577-.58 1.004-.865 1.155-.97zM14 17v2h4v-2h-4z"/>`),
		g.Group(children),
	)
}