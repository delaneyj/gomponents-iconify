package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSlashLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52.44 36a6 6 0 0 0-8.88 8l40.18 44.2c-.45.87-.9 1.75-1.32 2.64A62 62 0 1 0 72 214h88a85.23 85.23 0 0 0 32.35-6.3l11.21 12.3a6 6 0 0 0 8.88-8.08ZM160 202H72a50 50 0 1 1 5.9-99.64A86.25 86.25 0 0 0 74 128a6 6 0 0 0 12 0a73.92 73.92 0 0 1 6.44-30.2l91.22 100.34A73.65 73.65 0 0 1 160 202Zm86-74a85.85 85.85 0 0 1-21.85 57.27a6 6 0 0 1-4.47 2a6 6 0 0 1-4.47-10a74 74 0 0 0-99-108.92a6 6 0 1 1-7.11-9.67A86 86 0 0 1 246 128Z"/>`),
		g.Group(children),
	)
}