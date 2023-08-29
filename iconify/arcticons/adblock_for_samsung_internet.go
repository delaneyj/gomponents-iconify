package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdblockForSamsungInternet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.57 22.609l-.025-10.355h0c-.216-2.488-3.208-2.265-3.315 0l-.097 10.282"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.218 13.49c.062-1.89-3.383-2.132-3.448.36c0 0-.193 9.122-.072 9.098"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.549 14.044c0-1.981 3.214-2.239 3.214.024h0v11.637l1.62-3.387a2.444 2.444 0 0 1 2.154-1.5l1.693.048c.073.003.106.273.025.509l-4.017 11.589a6.136 6.136 0 0 1-3.58 3.725a8.087 8.087 0 0 1-7.235-.024c-2.67-1.434-3.524-3.945-3.725-6.363l.024-13.065a1.498 1.498 0 0 1 2.983-.236"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.663 5.5H16.337L5.5 16.337v15.326L16.337 42.5h15.326L42.5 31.663V16.337L31.663 5.5z"/>`),
		g.Group(children),
	)
}