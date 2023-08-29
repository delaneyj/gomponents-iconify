package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaleidoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.802" cy="24.486" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.927 16.949a11.306 11.306 0 1 1 5.134 21.38a11.428 11.428 0 0 1-1.512-.1M41.915 13.81a3.63 3.63 0 0 0-.164 1.008c0 1.586 1.687 3.383 1.097 4.753c-.51 1.186-3.375.83-4.239 1.765c-1.009 1.092-.72 3.926-2.089 4.554c-1.433.658-3.1-1.708-4.78-1.708c-1.47 0-3.45 2.236-4.739 1.725c-1.312-.52-.993-3.364-2.011-4.308c-.98-.908-3.308-.694-3.918-1.897c-.775-1.532.513-2.977.513-4.812c0-1.45-1.524-3.124-1.027-4.398c.574-1.469 3.243-2.141 4.339-3.237c.982-.982.347-3.582 1.635-4.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.849 20.815c0 1.095 1.164 2.335.757 3.28c-.352.819-2.33.573-2.925 1.218c-.697.754-.497 2.71-1.442 3.143c-.99.454-2.14-1.179-3.299-1.179c-1.015 0-2.381 1.544-3.27 1.191c-.906-.359-.686-2.321-1.388-2.973c-.676-.627-2.283-.479-2.704-1.31c-.535-1.057.354-2.054.354-3.32c0-1.001-1.052-2.156-.709-3.035c.396-1.014 1.792-1.329 2.548-2.085c.678-.678.884-2.224 1.773-2.617c.963-.425 2.425.778 3.545.778c1.147 0 2.286-1.142 3.267-.697c1.041.472.918 2.158 1.655 3.014c.506.587 2.167.606 2.48 1.327c.417.954-.642 2.157-.642 3.265Z"/><circle cx="26.642" cy="26.895" r="14.458" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}