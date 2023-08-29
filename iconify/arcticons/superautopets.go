package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superautopets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.8" cy="21.181" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M9.256 23.161c1.996 2.033 2.02 2.04 4.105.029"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 14.878c4.005-10.5 26.132-8.976 27.45 6.837"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.35 27.078a14.536 14.536 0 0 0 4.833 3.706m4.049 1.757a22.558 22.558 0 0 0 7.324 1.809m6.254-.47a19.552 19.552 0 0 0 6.953-2.445m4.11-3.36a20.972 20.972 0 0 0 3.627-5.68c-2.68.967-2.787.048-3.05-.68m-20.223-.072c.28 2.894-.777 4.385 1.142 5.912c5.51 3.443 19.1-.52 19.081-5.84m-34.1 5.363a6.562 6.562 0 0 1-1.85-4.755a7.618 7.618 0 0 1 7.878-7.33c4.35 0 7.474 2.62 7.85 6.65m-9.045 9.141a15.526 15.526 0 0 0-1.147 5.271c-.054.968 4.846 1.296 5.088.491a9.912 9.912 0 0 0 .108-4.005m7.324 1.809c-.388 2.864-.358 3.213-.04 4.527c.241.999 6.18 1.37 6.649.289c.57-1.31.116-4.21-.355-5.286"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.873 28.076c-.065.065.818 4.514.029 5.146c-1.066.854-2.94 1.552-3.932 1.214c-.812-.276-.22-3.015-.22-3.015M16.845 10.484L18.55 13.4l5.806 2.222l8.595-.59l1.909-3.256M18.55 13.4l-2.558 2.273m8.363-.051l.1 12.97m8.495-13.559l5.12 9.962"/><circle cx="8.767" cy="21.181" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}