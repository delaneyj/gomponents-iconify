package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3 10c0-3.771 0-5.657 1.172-6.828C5.343 2 7.229 2 11 2h2c3.771 0 5.657 0 6.828 1.172C21 4.343 21 6.229 21 10v4c0 3.771 0 5.657-1.172 6.828C18.657 22 16.771 22 13 22h-2c-3.771 0-5.657 0-6.828-1.172C3 19.657 3 17.771 3 14v-4Z" opacity=".5"/><path d="M7.25 6A.75.75 0 0 1 8 5.25h8a.75.75 0 0 1 0 1.5H8A.75.75 0 0 1 7.25 6Z"/><path fill-rule="evenodd" d="M17.714 14.649a5.75 5.75 0 1 0-.07.461l.09-.448l-.02-.014Zm-9.599-.857a17.778 17.778 0 0 0-.344-.218A4.25 4.25 0 0 1 16.25 14v.005c-.133-.023-.27-.045-.404-.065l-.015-.003a22.28 22.28 0 0 1-.486-.079c-.7-.126-1.267-.075-1.758.106c-.457.168-.801.436-1.055.633l-.024.018c-.275.214-.455.347-.678.43c-.203.074-.493.123-.973.036c-.704-.126-1.333-.426-1.915-.769c-.248-.147-.48-.296-.707-.443l-.12-.077Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}