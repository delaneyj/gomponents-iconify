package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatingIce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M4 4h64v64H4z"/><path fill="#61b2e4" d="m9.55 36.774l-.542.837a6.055 6.055 0 0 0-.227 6.205l4.315 7.868a6.054 6.054 0 0 0 1.68 1.935l5.594 4.188a6.056 6.056 0 0 0 3.478 1.206l16.985.422a6.054 6.054 0 0 0 1.878-.25l11.445-3.41a6.055 6.055 0 0 0 3.527-2.796l5.339-9.334a6.055 6.055 0 0 0-1.605-7.836"/><path fill="#fff" d="M19.705 25.833L9.551 32.48l4.175 8.353l8.191 6.132h19.71l13.227-4.578l6.563-11.579l-13-7.975l-18.25 5.75l-10.462-2.75z"/><path fill="#fff" d="M9.551 32.48v7l4.175 8.353l8.191 6.132h19.71l13.227-4.578l6.563-11.579v-7"/><path fill="#fff" d="m32.917 53.965l1.125-3.007l-1.5-1.375l1.062-2.618m-3.437-18.382l4.187 5.313l-3.187 3.687l1.093 3.25l-1.78 2.125a32.026 32.026 0 0 0 1.124 4.007"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.705 25.833L9.551 32.48l4.175 8.353l8.191 6.132h19.71l13.227-4.578l6.563-11.579l-13-7.975l-18.25 5.75l-10.462-2.75z"/><path d="M9.551 32.48v7l4.175 8.353l8.191 6.132h19.71l13.227-4.578l6.563-11.579v-7"/></g>`),
		g.Group(children),
	)
}