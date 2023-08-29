package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MachineWashPermanentPress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M66 298q4 22 21.5 36.5T127 349h274q22 0 39.5-14.5T462 298l58-264q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 83q-1-1-4.5-3.5T450 100q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q86 76 61 72L51 25q-1-8-9.5-13.5T25 8q-8 1-13.5 9.5T8 34zm12-168q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q17 17 32 23l-30 137q-3 17-21 17H127q-18 0-21-17L72 123q4 4 6 7zm335 283q22 0 22-21t-22-21H115q-22 0-22 21t22 21h298z"/>`),
		g.Group(children),
	)
}