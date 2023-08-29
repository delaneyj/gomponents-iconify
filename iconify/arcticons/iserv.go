package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iserv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.456 21.472l-2 5.3l-2-5.3m-4.159 2.15c0-1.1.9-2 2-2m-2 0v5.3" class="c"/><g class="c"><circle cx="13.296" cy="19.172" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.296 21.572v5.3"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.656 26.472c.4.3.8.4 1.6.4h.4c.7 0 1.3-.6 1.3-1.3s-.6-1.3-1.3-1.3h-.9c-.7 0-1.3-.6-1.3-1.3s.6-1.3 1.3-1.3h.4c.9 0 1.3.1 1.6.4" class="c"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.826 25.772c-.3.6-1 1-1.7 1c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2s2 .9 2 2v.7h-4m-.499 4.628l14.83.056"/><circle cx="13.293" cy="28.828" r=".75" fill="currentColor"/><circle cx="15.817" cy="28.828" r=".75" fill="currentColor"/><circle cx="18.34" cy="28.828" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}