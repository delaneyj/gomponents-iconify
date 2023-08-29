package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPickerCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopColorPickerCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M10.536 18.236c-1.56.617-4.257 2.372-4.728 1.9c-.306-.305.323-1.546.972-2.805h4.661l-.905.905Z" clip-rule="evenodd"/><path d="M18.803 8.555a1 1 0 1 1 1.415 1.414l-8.975 8.975a1 1 0 0 1-.34.223c-.396.156-.8.354-1.753.845c-.47.242-.679.348-.933.474c-1.748.861-2.405 1.068-3.116.357c-.71-.71-.504-1.367.351-3.105c.127-.257.234-.468.48-.945c.491-.953.689-1.356.845-1.753A1 1 0 0 1 7 14.701l8.975-8.975a1 1 0 1 1 1.414 1.414l-8.832 8.832c-.174.415-.393.856-.847 1.737a52.348 52.348 0 0 0-.535 1.06l.158-.077c.24-.119.441-.221.902-.458c.88-.454 1.321-.673 1.736-.847l8.832-8.832Z"/><path d="M17.39 7.14a1 1 0 1 1-1.415-1.414a3 3 0 1 1 4.243 4.243a1 1 0 1 1-1.415-1.414A1 1 0 0 0 17.39 7.14Zm-4.953.705a1 1 0 0 1 1.415-1.414l5.656 5.657a1 1 0 0 1-1.414 1.414l-5.657-5.657Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopColorPickerCircleFilled0)"/></g>`),
		g.Group(children),
	)
}