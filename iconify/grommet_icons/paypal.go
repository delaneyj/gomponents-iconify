package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.495 7.054c-1.07 4.973-4.48 7.604-9.89 7.604H9.643L8.18 24h3.182c.46 0 .85-.334.923-.788l.037-.198l.732-4.636l.047-.256a.933.933 0 0 1 .922-.788h.581c3.76 0 6.705-1.528 7.565-5.946c.345-1.773.179-3.26-.674-4.334M19.317 1.81C18.206.543 16.197 0 13.627 0H6.169c-.526 0-.973.383-1.055.9L2.008 20.598a.64.64 0 0 0 .633.74h4.604l1.157-7.335l-.036.23c.082-.518.526-.9 1.051-.9h2.188c4.299 0 7.664-1.747 8.648-6.797c.029-.15.076-.438.076-.438c.279-1.869-.002-3.137-1.012-4.287"/>`),
		g.Group(children),
	)
}