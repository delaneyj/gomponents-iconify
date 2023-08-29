package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAntarctica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23.59 33.2c.167-.868-.043-1.763-.564-1.353c-1.067-.485-2.39-2.864-2.527-3.269c4.986 3.782 9.793 4.234 10.72-1.684c.925-5.918 4.905-4.075 7.306-3.63c2.983.484 6.239 1.854 8.872 3.304c1.459.763 1.705 3.005 1.34 4.707c-.786 1.686.663 2.657 1.79 3.562c1.518 1.584.956 3.955.427 5.831c-.498 1.562-2.163 3.021-1.797 4.634c.199.875-.901 1.524-1.493 2.192c-2.176 2.243-5.57 2.092-8.394 1.502c-1.493.024.52-.807.176-1.801c.63-1.816-1.043-2.689-2.85-2.888c-2.537-.314-3.006-.025-5.543-.6c-1.324.893-5.073-.8-5.456-2.743c-.323-1.639-1.8-2.888-2.05-4.516c1.44.282-.25-1.575 1.29-2.164c-.08-.766-.705-1.763-1.246-1.084z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}