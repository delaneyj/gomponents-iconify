package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liftago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 34.588a49.601 49.601 0 0 1 15.78-2.616c1.631 3.105 5.255 4.065 9.327 4.079a49.16 49.16 0 0 0 12.07-1.818L43.5 13.414a48.458 48.458 0 0 1-15.87 2.825c-1.175-1.984-4.681-4.29-11.497-4.29A41.297 41.297 0 0 0 4.5 14.032"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.692 19.058v7.714h3.857m20.33-5.11v5.785a1.929 1.929 0 0 1-1.929 1.929h0a1.923 1.923 0 0 1-1.364-.565"/><rect width="3.857" height="5.111" x="28.021" y="21.662" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.929" transform="rotate(-180 29.95 24.217)"/><rect width="3.857" height="5.111" x="33.835" y="21.662" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.929"/><circle cx="13.424" cy="19.299" r=".675" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.424 21.662v5.11m12.586-1.928a1.929 1.929 0 0 1-1.929 1.928h0a1.929 1.929 0 0 1-1.928-1.928V23.59a1.929 1.929 0 0 1 1.928-1.929h0a1.929 1.929 0 0 1 1.929 1.929m0 3.182v-5.11m-6.473-1.592v5.738a.964.964 0 0 0 .964.964h.29m-4.414.001v-6.366a1.35 1.35 0 0 1 1.35-1.35h0a2.63 2.63 0 0 1 .436.034m-2.917 2.571h5.304"/>`),
		g.Group(children),
	)
}