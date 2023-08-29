package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AoOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="red" d="M0 0h512v259.8H0z"/><path d="M0 252.2h512V512H0z"/></g><path fill="#ffec00" fill-rule="evenodd" d="M228.7 148.2c165.2 43.3 59 255.6-71.3 167.2l-8.8 13.6c76.7 54.6 152.6 10.6 174-46.4c22.2-58.8-7.6-141.5-92.6-150l-1.3 15.6z"/><path fill="#ffec00" fill-rule="evenodd" d="m170 330.8l21.7 10.1l-10.2 21.8l-21.7-10.2zm149-99.5h24v24h-24zm-11.7-38.9l22.3-8.6l8.7 22.3l-22.3 8.7zm-26-29.1l17.1-16.9l16.9 17l-17 16.9zm-26.2-39.8l22.4 8.4l-8.5 22.4l-22.4-8.4zM316 270l22.3 8.9l-9 22.2l-22.2-8.9zm-69.9 70l22-9.3l9.5 22l-22 9.4zm-39.5 2.8h24v24h-24zm41.3-116l-20.3-15l-20.3 14.6l8-23l-20.3-15h24.5l8.5-22.6l7.8 22.7l24.7-.3l-19.6 15.3l7 23.4z"/><path fill="#fe0" fill-rule="evenodd" d="M336 346.4c-1.2.4-6.2 12.4-9.7 18.2l3.7 1c13.6 4.8 20.4 9.2 26.2 17.5a7.9 7.9 0 0 0 10.2.7s2.8-1 6.4-5c3-4.5 2.2-8-1.4-11.1c-11-8-22.9-14-35.4-21.3z"/><path fill-rule="evenodd" d="M365.3 372.8a4.3 4.3 0 1 1-8.7 0a4.3 4.3 0 0 1 8.6 0zm-21.4-13.6a4.3 4.3 0 1 1-8.7 0a4.3 4.3 0 0 1 8.7 0zm10.9 7a4.3 4.3 0 1 1-8.7 0a4.3 4.3 0 0 1 8.7 0z"/><path fill="#fe0" fill-rule="evenodd" d="M324.5 363.7c-42.6-24.3-87.3-50.5-130-74.8c-18.7-11.7-19.6-33.4-7-49.9c1.2-2.3 2.8-1.8 3.4-.5c1.5 8 6 16.3 11.4 21.5A5288 5288 0 0 1 334 345.6c-3.4 5.8-6 12.3-9.5 18z"/><path fill="#ffec00" fill-rule="evenodd" d="m297.2 305.5l17.8 16l-16 17.8l-17.8-16z"/><path fill="none" stroke="#000" stroke-width="3" d="m331.5 348.8l-125-75.5m109.6 58.1L274 304.1m18.2 42.7L249.3 322"/>`),
		g.Group(children),
	)
}