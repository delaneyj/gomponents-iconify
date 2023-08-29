package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YapiKredi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.125 20.905c.39.654 1.36.654 1.75 0c1.663-2.798 5.052-6.486 10.103-6.486c6.217 0 9.522 3.384 9.522 8.578s-3.227 8.263-6.61 8.263c-3.384 0-5.706-1.968-4.84-6.453c0 0 .612 2.544 3.541 2.4c3.994-.197 5.351-7.83-1.298-7.83c-8.21 0-5.01 14.204-11.293 14.204c-6.283 0-3.082-14.204-11.293-14.204c-6.65 0-5.292 7.633-1.298 7.83c2.93.144 3.54-2.4 3.54-2.4c.866 4.485-1.455 6.453-4.839 6.453S3.5 28.19 3.5 22.997s3.305-8.578 9.522-8.578c5.051 0 8.44 3.688 10.103 6.486Z"/>`),
		g.Group(children),
	)
}