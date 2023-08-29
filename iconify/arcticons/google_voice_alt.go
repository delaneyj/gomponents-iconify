package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleVoiceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A16.24 16.24 0 0 0 7.76 20.74h0c0 8.51 7.74 15.58 16.24 16.18v6.58c8-2.23 16.24-12.47 16.24-22.76S33 4.5 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.575 24.248a30.695 30.695 0 0 1-3.118-.55a1.54 1.54 0 0 0-1.485.288a236.71 236.71 0 0 1-1.807 1.787a14.201 14.201 0 0 1-6.484-6.49c.767-.787 1.485-1.539 1.777-1.801a1.54 1.54 0 0 0 .287-1.485a30.598 30.598 0 0 1-.55-3.124a.99.99 0 0 0-1.086-.883a.994.994 0 0 0-.052.007H15.84a.743.743 0 0 0-.673.678c-.272 3.807 1.851 7.88 2.282 8.663v.03l.06.113h0a17.541 17.541 0 0 0 6.434 6.435h0l.218.124h0c.99.525 4.925 2.505 8.602 2.232a.743.743 0 0 0 .689-.673v-4.212a.99.99 0 0 0-.825-1.132a.91.91 0 0 0-.052-.007Z"/>`),
		g.Group(children),
	)
}