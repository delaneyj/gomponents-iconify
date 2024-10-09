package marketeq

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 50 50")
)

func Acorn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m37.5 20.833l-.937 8.542a17.42 17.42 0 0 1-10.792 14.23v0a2.08 2.08 0 0 1-1.542 0v0a17.42 17.42 0 0 1-10.791-14.23l-.938-8.542zM27.083 6.25A10.2 10.2 0 0 0 25 12.5"/><path stroke="#344054" d="M16.667 12.5h16.666a6.25 6.25 0 0 1 6.25 6.25a2.083 2.083 0 0 1-2.083 2.083h-25a2.083 2.083 0 0 1-2.083-2.083a6.25 6.25 0 0 1 6.25-6.25"/></g>`), g.Group(children),
	)
}

func AcrobaticTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="m22.5 6.25l13.646 5.104a2.082 2.082 0 0 1 .75 3.438l-3.98 3.958l-5.833-1.98a2.4 2.4 0 0 0-.75-.103h-8.458a2.08 2.08 0 0 0-1.396.541L12.5 20.833"/><path stroke="#344054" stroke-width="3" d="m33.125 19.458l.375 1.375a8.77 8.77 0 0 1-5.958 10.771v0"/><path stroke="#306CFE" stroke-width="2" d="M22.917 43.75H25a2.083 2.083 0 0 0 2.083-2.083v-12.5"/><path stroke="#306CFE" stroke-width="2" d="M22.917 33.333a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/></g>`), g.Group(children),
	)
}

func ActionCam(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M34.27 14.583h.21"/><path stroke="#344054" stroke-width="2" d="M14.583 16.667h2.084"/><path stroke="#306CFE" stroke-width="2" d="M29.167 31.25h-8.334v12.5h8.334z"/><path stroke="#306CFE" stroke-width="2" d="M41.667 22.917v6.25a2.083 2.083 0 0 1-2.084 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083v-18.75a2.083 2.083 0 0 1 2.083-2.084H25"/><path stroke="#344054" stroke-width="2" d="M43.75 20.833v-12.5c0-1.15-.933-2.083-2.083-2.083H27.083c-1.15 0-2.083.933-2.083 2.083v12.5c0 1.15.933 2.084 2.083 2.084h14.584c1.15 0 2.083-.933 2.083-2.084"/></g>`), g.Group(children),
	)
}

func ActionCamera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 33.333v6.25a2.084 2.084 0 0 1-2.084 2.084H8.333a2.083 2.083 0 0 1-2.083-2.084V12.5a2.083 2.083 0 0 1 2.083-2.083H18.75"/><path stroke="#344054" d="M43.75 31.25V10.417a2.083 2.083 0 0 0-2.083-2.084H20.833a2.083 2.083 0 0 0-2.083 2.084V31.25a2.083 2.083 0 0 0 2.083 2.083h20.834a2.083 2.083 0 0 0 2.083-2.083m-12.5-14.583a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func AdapterThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 31.25v12.5m12.5-33.333h6.25zm0 8.333h6.25z"/><path stroke="#306CFE" d="M16.667 6.25h12.5a2.083 2.083 0 0 1 2.083 2.083v20.834a2.083 2.083 0 0 1-2.083 2.083H14.583a2.083 2.083 0 0 1-2.083-2.083v-18.75a4.167 4.167 0 0 1 4.167-4.167"/></g>`), g.Group(children),
	)
}

func AddCollection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 14.583v12.5m-6.25-6.25h12.5z"/><path stroke="#306CFE" d="M14.583 43.75h27.084a2.083 2.083 0 0 0 2.083-2.083v-31.25"/><path stroke="#306CFE" d="M33.333 6.25h-25c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.084 2.083 2.084h25c1.15 0 2.084-.933 2.084-2.084v-25c0-1.15-.933-2.083-2.084-2.083"/></g>`), g.Group(children),
	)
}

func AddPlaylist(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 39.583H43.75M29.167 31.25H43.75zm6.25-20.833h8.333zm4.166-4.167v8.333z"/><path stroke="#306CFE" d="M13.542 43.75a7.292 7.292 0 1 0 0-14.583a7.292 7.292 0 0 0 0 14.583"/><path stroke="#306CFE" d="M29.167 20.833a5.96 5.96 0 0 0-2.084-8.333a48.3 48.3 0 0 1-6.25-6.25v30.208"/></g>`), g.Group(children),
	)
}

func AddPlaylistTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 39.583H43.75M29.167 31.25H43.75zm6.25-20.833h8.333zm4.166-4.167v8.333z"/><path stroke="#306CFE" d="M20.833 36.458V6.25zm-7.291-7.291a7.292 7.292 0 1 0 0 14.583a7.292 7.292 0 0 0 0-14.583"/></g>`), g.Group(children),
	)
}

func Agenda(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25v12.5l-4.167-2.083l-4.166 2.083V6.25zm18.75 8.333h-4.167v4.167h4.167zm0 12.5h-4.167v4.167h4.167z"/><path stroke="#306CFE" d="M39.583 41.667V8.333c0-1.15-.932-2.083-2.083-2.083H10.417c-1.15 0-2.084.933-2.084 2.083v33.334c0 1.15.933 2.083 2.084 2.083H37.5c1.15 0 2.083-.933 2.083-2.083"/></g>`), g.Group(children),
	)
}

func AgendaLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 25v14.583l-4.166 4.167l-4.167-4.167V25a2.083 2.083 0 0 1 2.083-2.083H12.5A2.083 2.083 0 0 1 14.583 25"/><path stroke="#306CFE" d="M43.75 33.333h-4.167m4.167-16.666h-4.167zm0 8.333h-4.167zM22.917 43.75h16.666a2.084 2.084 0 0 0 2.084-2.083V8.333a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083v6.25"/></g>`), g.Group(children),
	)
}

func Aim(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m35.417 43.75l-3.625-12.5zm-20.834 0l3.625-12.5zM25 8.333V6.25zm4.167 12.5a4.167 4.167 0 1 0-8.334 0a4.167 4.167 0 0 0 8.334 0"/><path stroke="#344054" d="M25 33.333c6.904 0 12.5-5.596 12.5-12.5c0-6.903-5.596-12.5-12.5-12.5s-12.5 5.597-12.5 12.5c0 6.904 5.596 12.5 12.5 12.5"/></g>`), g.Group(children),
	)
}

func AimTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 8.333V6.25zm-4.167 25V43.75zm8.334 0V43.75zm-12.5 10.417h16.666zm12.5-22.917a4.166 4.166 0 1 0-8.332 0a4.166 4.166 0 0 0 8.332 0"/><path stroke="#344054" d="M25 33.333c6.904 0 12.5-5.596 12.5-12.5c0-6.903-5.596-12.5-12.5-12.5s-12.5 5.597-12.5 12.5c0 6.904 5.596 12.5 12.5 12.5"/></g>`), g.Group(children),
	)
}

func Alpha(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.667 12.5s-7.438 27.083-22.813 27.083c-11.75 0-13.166-6.979-12.5-12.979c.834-6.583 6.875-16.187 16.021-16.187c14.583 0 9.063 29.166 21.375 29.166"/>`), g.Group(children),
	)
}

func AlphaCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M32.417 19.646S28.937 31.25 22.292 31.25c-4 0-5.98-2.292-5.521-5.333c.5-3.438 3.354-7.167 6.791-7.167c6.459 0 3.709 12.5 9.771 12.5"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M10.083 35.417h-1.75a2.083 2.083 0 0 1-2.083-2.084V12.5a2.083 2.083 0 0 1 2.083-2.083h26.084a2.08 2.08 0 0 1 1.625.791l7.25 9.063c.294.366.455.822.458 1.292v11.77a2.084 2.084 0 0 1-2.083 2.084h-1.73m-8.917 0H19"/><path stroke="#344054" d="M25 18.75v8.333m-6.25 8.334a4.166 4.166 0 1 1-8.332 0a4.166 4.166 0 0 1 8.332 0m20.833 0a4.166 4.166 0 1 1-8.332 0a4.166 4.166 0 0 1 8.332 0m-18.75-12.5h8.334z"/></g>`), g.Group(children),
	)
}

func AngelFortyFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 39.583a16.67 16.67 0 0 0-4.875-11.791"/><path stroke="#306CFE" d="M43.75 39.583H6.25l29.167-29.166"/></g>`), g.Group(children),
	)
}

func AppartmentTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 22.917h8.334m0 20.833v-12.5h-8.334v12.5zm-8.334-29.167h8.334z"/><path stroke="#306CFE" d="M31.25 43.75h-25V8.333A2.083 2.083 0 0 1 8.333 6.25h20.834a2.083 2.083 0 0 1 2.083 2.083zm0 0h12.5V18.167a2.084 2.084 0 0 0-1.417-2.084L31.25 12.5z"/></g>`), g.Group(children),
	)
}

func AppleFruit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583a10.65 10.65 0 0 1 4.167-8.333"/><path stroke="#306CFE" d="M41.5 25.792c-1.042 9.479-5.417 17.958-9.854 17.958c-2.896 0-3.563-2.083-6.646-2.083s-3.77 2.083-6.646 2.083c-4.437 0-8.812-8.48-9.854-17.958c-.73-6.73.958-10.917 5.042-12.625c5.208-2.188 8.146 1.416 11.458 1.416c3.313 0 6.25-3.604 11.438-1.416c4.083 1.708 5.812 5.895 5.062 12.625"/></g>`), g.Group(children),
	)
}

func ArmchairFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 18.75v-4.167a8.333 8.333 0 0 1 8.333-8.333h8.334a8.333 8.333 0 0 1 8.333 8.333v4.167m-2.083 25v-4.167M14.583 43.75v-4.167z"/><path stroke="#306CFE" d="M37.167 18.75a6.52 6.52 0 0 0-5.917 6.604v5.896h-12.5v-5.896a6.52 6.52 0 0 0-5.917-6.604a6.25 6.25 0 0 0-2.416 12.125V37.5a2.083 2.083 0 0 0 2.083 2.083h25a2.083 2.083 0 0 0 2.083-2.083v-6.625a6.25 6.25 0 0 0-2.416-12.125"/></g>`), g.Group(children),
	)
}

func Arrival(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 43.75h37.5"/><path stroke="#306CFE" d="m41.146 28.77l-.5 1.272a1.896 1.896 0 0 1-2.48 1.083L27.75 27.25a1.87 1.87 0 0 0-1.562 0l-6.48 3.042c-.401.192-.854.25-1.291.166L11.5 29l6.5-5.396l-8.125-3.041a2.46 2.46 0 0 1-1.542-2.355L8.854 6.25l4.48 1.083a2.33 2.33 0 0 1 1.75 1.938l.604 3.52l19.291 4.667a8.606 8.606 0 0 1 6.25 11.313z"/></g>`), g.Group(children),
	)
}

func Atom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 25h.208"/><path stroke="#306CFE" stroke-width="2" d="M33.333 25c0 10.417-3.729 18.75-8.333 18.75S16.667 35.417 16.667 25S20.396 6.25 25 6.25s8.333 8.333 8.333 18.75M8.854 34.896c2.292 4.166 11.375 3.208 20.292-2.271c8.917-5.48 14.27-13.312 12-17.52c-2.271-4.21-11.375-3.21-20.313 2.27S6.583 30.688 8.854 34.895m32.292 0c2.27-4.167-3.084-12.042-12-17.521s-18-6.5-20.292-2.27c-2.291 4.228 3.084 11.978 11.98 17.52c8.895 5.542 18.02 6.5 20.312 2.27"/></g>`), g.Group(children),
	)
}

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M37.5 20.833V31.25a12.5 12.5 0 0 1-25 0V14.583a8.333 8.333 0 0 1 8.333-8.333v0a8.333 8.333 0 0 1 8.334 8.333V31.25a4.167 4.167 0 1 1-8.334 0V14.583"/>`), g.Group(children),
	)
}

func Avocado(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 6.25A10.2 10.2 0 0 0 25 12.5"/><path stroke="#306CFE" d="M34.396 28.125a26.5 26.5 0 0 1-2.688-8.333a7.79 7.79 0 0 0-3.208-6.25a6.4 6.4 0 0 0-7 0a7.8 7.8 0 0 0-3.208 6.25a26.5 26.5 0 0 1-2.688 8.333a16 16 0 0 0-1.02 6.375A9.896 9.896 0 0 0 25 43.75a9.896 9.896 0 0 0 10.417-9.25a16 16 0 0 0-1.021-6.375"/></g>`), g.Group(children),
	)
}

func AvocadoTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 35.417a5.56 5.56 0 0 1-5.208-5.834c0-3.208 2.333-8.75 5.208-8.75s5.208 5.542 5.208 8.75A5.563 5.563 0 0 1 25 35.417"/><path stroke="#306CFE" d="M35.167 15.625C32.729 10.521 29.937 6.25 25 6.25s-7.73 4.27-10.167 9.375C12.396 20.729 10.417 25 10.417 31.25c0 6.896 6.541 12.5 14.583 12.5s14.583-5.604 14.583-12.5c0-6.25-2.083-10.667-4.416-15.625"/></g>`), g.Group(children),
	)
}

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 25a22.9 22.9 0 0 1 12.5 6.25a19.7 19.7 0 0 0 4.167-12.5a19.7 19.7 0 0 0-4.167-12.5a22.9 22.9 0 0 1-12.5 6.25z"/><path stroke="#306CFE" d="M18.75 43.75h-4.167a2.083 2.083 0 0 1-2.083-2.083v-31.25a4.167 4.167 0 1 1 8.333 0v31.25a2.083 2.083 0 0 1-2.083 2.083"/></g>`), g.Group(children),
	)
}

func BabyCribTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 25V12.5m0 20.833h2.084zm18.75 0h2.084zM25 12.5V25zm10.417 0V25z"/><path stroke="#306CFE" d="M6.25 12.5V8.333m0 33.334h37.5V25H6.25zm0 0H25V25H6.25zM43.75 25V12.5H6.25V25zm0-12.5V8.333z"/></g>`), g.Group(children),
	)
}

func Backpack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 20.833h12.5zm12.5 12.5a2.083 2.083 0 0 0-2.083-2.083h-8.334a2.083 2.083 0 0 0-2.083 2.083V43.75h12.5zm-2.083-25a2.083 2.083 0 0 0-2.084-2.083h-4.166a2.083 2.083 0 0 0-2.084 2.083v2.084h8.334zm10.416 31.25h2.084A2.083 2.083 0 0 0 43.75 37.5v-6.25h-4.167zM10.417 31.25H6.25v6.25a2.083 2.083 0 0 0 2.083 2.083h2.084z"/><path stroke="#306CFE" d="M18.75 10.417h12.5a8.333 8.333 0 0 1 8.333 8.333v22.917A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083V18.75a8.333 8.333 0 0 1 8.333-8.333"/></g>`), g.Group(children),
	)
}

func BackwardTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 6.25v37.5"/><path stroke="#306CFE" d="m18.75 25l20.833 14.583V10.417z"/></g>`), g.Group(children),
	)
}

func Badge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.708 27.917l-2.875.395l2.084 1.855l-.5 2.645L25 31.563l2.583 1.25l-.5-2.645l2.084-1.855l-2.875-.395L25 25.52z"/><path stroke="#306CFE" d="M25 41.667a12.5 12.5 0 1 1 0-24.999a12.5 12.5 0 0 1 0 24.999m0-25a12.5 12.5 0 0 1 6.25 1.687V6.25h-12.5v12.104A12.5 12.5 0 0 1 25 16.667"/></g>`), g.Group(children),
	)
}

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.333 43.75H12.667a2.084 2.084 0 0 1-2.084-2.25l1.771-22.917a2.084 2.084 0 0 1 2.084-1.916h21.124a2.083 2.083 0 0 1 2.084 1.916l1.77 22.917a2.084 2.084 0 0 1-2.083 2.25"/><path stroke="#344054" d="M18.75 22.917V12.5A6.25 6.25 0 0 1 25 6.25v0a6.25 6.25 0 0 1 6.25 6.25v10.417"/></g>`), g.Group(children),
	)
}

func BagAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 14.583v4.167a6.25 6.25 0 0 0 12.5 0v-4.167"/><path stroke="#306CFE" d="M41.417 43.75H8.583A2.082 2.082 0 0 1 6.5 41.438l3.708-33.334a2.083 2.083 0 0 1 2.084-1.854h25.416a2.083 2.083 0 0 1 2.084 1.854L43.5 41.438a2.084 2.084 0 0 1-2.083 2.312"/></g>`), g.Group(children),
	)
}

func BagAltOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 35.417v-8.334M25 35.417v-8.334zm-8.333 0v-8.334z"/><path stroke="#306CFE" d="m31.25 6.25l4.167 12.5m2.395 25H12.189A2.08 2.08 0 0 1 10.104 42L6.25 18.75h37.5L39.875 42a2.08 2.08 0 0 1-2.062 1.75M18.75 6.25l-4.167 12.5z"/></g>`), g.Group(children),
	)
}

func BananaLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.333 12.5V8.333a2.083 2.083 0 0 1 2.084-2.083h4.166a2.083 2.083 0 0 1 2.084 2.083V12.5A18.94 18.94 0 0 0 25 29.333c4.167 2.73 10.417 3.271 14.583 3.313a4.167 4.167 0 0 1 4.167 4.166a4.17 4.17 0 0 1-2.417 3.771c-5.604 2.5-19.708 7.167-29.896-3.458C2.083 27.25 8.333 12.5 8.333 12.5"/>`), g.Group(children),
	)
}

func BananasLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M28.667 31.083a30.1 30.1 0 0 0 13.125-2.958a3.458 3.458 0 0 0-2.334-6.458c-2.975.691-6.037.93-9.083.708c-7.02-.604-11.48-5.542-13.52-7.792"/><path stroke="#306CFE" d="M8.333 12.5V8.333a2.083 2.083 0 0 1 2.084-2.083h4.166a2.083 2.083 0 0 1 2.084 2.083V12.5A18.94 18.94 0 0 0 25 29.333c4.167 2.73 10.417 3.271 14.583 3.313a4.167 4.167 0 0 1 4.167 4.166a4.17 4.17 0 0 1-2.417 3.771c-5.604 2.5-19.708 7.167-29.896-3.458C2.083 27.25 8.333 12.5 8.333 12.5"/></g>`), g.Group(children),
	)
}

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 18.75v25m-12.5-25v25zm25 0v25z"/><path stroke="#306CFE" d="M6.25 43.75h37.5m-37.5-25L25 6.25l18.75 12.5z"/></g>`), g.Group(children),
	)
}

func BarcodeScan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 35.417v-2.084M8.333 25h33.334zm8.334-10.417v2.084zm8.333 0v2.084zm-8.333 20.834v-2.084zm16.666-20.834v2.084zm0 20.834v-2.084z"/><path stroke="#306CFE" d="M6.25 16.667V8.333A2.083 2.083 0 0 1 8.333 6.25h8.334M43.75 16.667V8.333a2.083 2.083 0 0 0-2.083-2.083h-8.334M6.25 33.333v8.334a2.083 2.083 0 0 0 2.083 2.083h8.334m16.666 0h8.334a2.083 2.083 0 0 0 2.083-2.083v-8.334"/></g>`), g.Group(children),
	)
}

func Barrier(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 43.75h8.334M12.5 10.417V6.25zm0 25v8.333zm25-25V6.25zm-25 16.666V18.75zm25 0V18.75zm0 8.334v8.333zm-4.167 8.333h8.334z"/><path stroke="#306CFE" d="M41.667 18.75H8.333a2.083 2.083 0 0 1-2.083-2.083V12.5a2.083 2.083 0 0 1 2.083-2.083h33.334A2.083 2.083 0 0 1 43.75 12.5v4.167a2.083 2.083 0 0 1-2.083 2.083m0 16.667H8.333a2.083 2.083 0 0 1-2.083-2.084v-4.166a2.083 2.083 0 0 1 2.083-2.084h33.334a2.083 2.083 0 0 1 2.083 2.084v4.166a2.083 2.083 0 0 1-2.083 2.084"/></g>`), g.Group(children),
	)
}

func BasketballHoop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m10.417 16.667l2.083 25M12.5 25h25zm25 8.333h-25zm2.083-16.666l-2.083 25zm-9.729 0l-.687 25zm-9.708 0l.687 25z"/><path stroke="#344054" d="M39.583 8.333H10.417a4.167 4.167 0 1 0 0 8.334h29.166a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func BathroomTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M33.333 12.5v-2.083a4.167 4.167 0 0 0-4.166-4.167H18.75a4.167 4.167 0 0 0-4.167 4.167V43.75"/><path stroke="#306CFE" d="M39.583 20.833h-12.5V18.75a6.25 6.25 0 0 1 6.25-6.25v0a6.25 6.25 0 0 1 6.25 6.25z"/><path stroke="#344054" d="M29.167 33.333v-4.166M10.417 43.75h8.333zM37.5 33.333v-4.166z"/></g>`), g.Group(children),
	)
}

func BathtubEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 20.833V12.5a4.167 4.167 0 0 0-8.333 0m5.208 25l1.042 4.167M13.542 37.5L12.5 41.667z"/><path stroke="#306CFE" d="M43.75 20.833H6.25zM16.667 37.5h16.666a8.333 8.333 0 0 0 8.334-8.333v-8.334H8.333v8.334a8.334 8.334 0 0 0 8.334 8.333"/></g>`), g.Group(children),
	)
}

func BathtubThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 14.583a4.167 4.167 0 0 0-4.167-4.166v0a4.167 4.167 0 0 0-4.166 4.166v8.334"/><path stroke="#306CFE" d="M6.25 22.917h37.5zm35.417 8.333v-8.333H8.333v8.333a8.333 8.333 0 0 0 8.334 8.333h16.666a8.333 8.333 0 0 0 8.334-8.333"/></g>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29.167 10.417h-8.334V6.25h8.334zm6.25 31.25V12.5a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083v29.167a2.083 2.083 0 0 0 2.084 2.083h16.666a2.083 2.083 0 0 0 2.084-2.083"/>`), g.Group(children),
	)
}

func BatteryCharge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m25 31.25l2.083-4.167h-4.166L25 22.917"/><path stroke="#306CFE" d="M35.417 41.667V12.5a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083v29.167a2.083 2.083 0 0 0 2.084 2.083h16.666a2.083 2.083 0 0 0 2.084-2.083M29.167 6.25h-8.334v4.167h8.334z"/></g>`), g.Group(children),
	)
}

func BatteryOneHundred(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 18.75h-4.166v16.667h4.166z"/><path stroke="#306CFE" d="M35.417 41.667V12.5a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083v29.167a2.083 2.083 0 0 0 2.084 2.083h16.666a2.083 2.083 0 0 0 2.084-2.083m-6.25-31.25h-8.334V6.25h8.334z"/></g>`), g.Group(children),
	)
}

func BatteryOneHundredLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 18.75h-4.166m4.166 16.667h-4.166zm0-8.334h-4.166z"/><path stroke="#306CFE" d="M35.417 41.667V12.5a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083v29.167a2.083 2.083 0 0 0 2.084 2.083h16.666a2.083 2.083 0 0 0 2.084-2.083M29.167 6.25h-8.334v4.167h8.334z"/></g>`), g.Group(children),
	)
}

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 37.5v-4.167m6.25-18.75A2.083 2.083 0 0 0 12.5 12.5H8.333a2.083 2.083 0 0 0-2.083 2.083V25h8.333zM41.667 37.5v-4.167z"/><path stroke="#306CFE" d="M6.25 25h35.417a2.083 2.083 0 0 1 2.083 2.083v6.25H6.25z"/></g>`), g.Group(children),
	)
}

func BedroomFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M31.25 22.917h12.5v12.5H6.25v-12.5m37.5 16.666V10.417M6.25 39.583v-6.25z"/><path stroke="#344054" d="M31.25 27.083h-25v-6.25a2.083 2.083 0 0 1 2.083-2.083h20.834a2.083 2.083 0 0 1 2.083 2.083z"/></g>`), g.Group(children),
	)
}

func BedroomSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 35.417v4.166m-29.166-4.166v4.166zm0-20.834v-4.166m20.833 4.166a4.167 4.167 0 0 0-4.167-4.166h-4.166a4.167 4.167 0 0 0-4.167 4.166h-8.333v12.5h29.166v-12.5zm8.333 0v-4.166z"/><path stroke="#306CFE" d="M10.417 27.083h29.166a4.167 4.167 0 0 1 4.167 4.167v4.167H6.25V31.25a4.167 4.167 0 0 1 4.167-4.167"/></g>`), g.Group(children),
	)
}

func BenchPress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="M41.667 33.333v8.334M16.667 25v16.667zM6.25 33.333h37.5zm2.083 0v8.334z"/><path stroke="#344054" stroke-width="2" d="M16.667 25a8.333 8.333 0 1 0 0-16.667a8.333 8.333 0 0 0 0 16.667"/><path stroke="#344054" stroke-width="3" d="M16.563 16.667h.208"/></g>`), g.Group(children),
	)
}

func BikeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 33.333a6.25 6.25 0 1 1-12.5 0a6.25 6.25 0 0 1 12.5 0m-31.25-6.25a6.25 6.25 0 1 0 0 12.5a6.25 6.25 0 0 0 0-12.5"/><path stroke="#306CFE" d="M29.167 10.417h2.416a2.08 2.08 0 0 1 2.084 1.708L37.5 33.333"/><path stroke="#306CFE" d="M12.5 33.333h6.25l15.73-16.666"/><path stroke="#306CFE" d="M16.667 16.667h3.291a2.084 2.084 0 0 1 2.084 1.666l1.812 9.105"/></g>`), g.Group(children),
	)
}

func Bikecycle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 16.667H31.77a22.02 22.02 0 0 0-21.354 16.666v0"/><path stroke="#344054" d="M6.25 37.5a4.167 4.167 0 1 0 8.333 0a4.167 4.167 0 0 0-8.333 0m20.833-4.167a8.333 8.333 0 1 0 16.667 0a8.333 8.333 0 0 0-16.667 0"/><path stroke="#306CFE" d="M18.75 20.833v-4.166a2.083 2.083 0 0 0-2.083-2.084h-2.084m20.834 18.75V10.417a2.083 2.083 0 0 0-2.084-2.084h-4.166"/></g>`), g.Group(children),
	)
}

func BillDollar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 27.083h5.209a3.126 3.126 0 0 0 0-6.25h-2.084a3.125 3.125 0 0 1 0-6.25h5.209M25 22.917h-8.333M37.5 14.583V12.5zm0 14.584v-2.084zM25 31.25h-8.333z"/><path stroke="#306CFE" d="M29.167 6.25h-18.75a2.083 2.083 0 0 0-2.084 2.083V43.75l4.875-2.083l4.855 2.083l4.854-2.083l4.854 2.083l4.854-2.083L37.5 43.75V37.5"/></g>`), g.Group(children),
	)
}

func BilliardBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 22.917a6.25 6.25 0 1 0 0 12.499a6.25 6.25 0 0 0 0-12.5m0-8.334a4.166 4.166 0 1 0 0 8.333a4.166 4.166 0 0 0 0-8.333"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func BirthdayCake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 18.75v8.333M25 6.25s-4.167 4.875-4.167 8.333a4.167 4.167 0 1 0 8.334 0C29.167 11.125 25 6.25 25 6.25"/><path stroke="#306CFE" d="M6.25 43.75h37.5m-2.083 0H8.333V29.167a2.084 2.084 0 0 1 2.084-2.084h29.166a2.083 2.083 0 0 1 2.084 2.084z"/></g>`), g.Group(children),
	)
}

func BlackBoard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 35.417v6.25m12.5 2.083l-1.396-8.333l-.354-2.084zM31.583 8.333L31.25 6.25zm-17.333 25l-.354 2.084L12.5 43.75zm4.5-27.083l-.354 2.083z"/><path stroke="#306CFE" d="M41.667 8.333H8.333c-1.15 0-2.083.933-2.083 2.084V31.25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083V10.417c0-1.15-.933-2.084-2.083-2.084"/></g>`), g.Group(children),
	)
}

func BlueprintArchitecture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 18.75h-4.584l-4.166 4.167h-5.834V18.75h-12.5v16.667h-4.166a4.167 4.167 0 1 0 0 8.333h31.25a2.083 2.083 0 0 0 2.083-2.083V20.833a2.083 2.083 0 0 0-2.083-2.083M14.583 12.5v22.917h-4.166a4.167 4.167 0 0 0-4.167 4.166V16.667a4.167 4.167 0 0 1 4.167-4.167z"/><path stroke="#344054" d="m43.146 9.77l-2.917-2.916a2.083 2.083 0 0 0-2.916 0l-10.23 10.23v5.833h5.834l10.229-10.23a2.083 2.083 0 0 0 0-2.916"/></g>`), g.Group(children),
	)
}

func BluntedCone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 43.75c9.205 0 16.667-2.798 16.667-6.25S34.205 31.25 25 31.25S8.333 34.048 8.333 37.5S15.795 43.75 25 43.75"/><path stroke="#306CFE" d="M41.667 36.98L35.417 10C34.896 7.917 30.458 6.25 25 6.25S15.167 7.896 14.583 10l-6.25 27.083a2 2 0 0 0 0 .417c0 3.458 7.459 6.25 16.667 6.25s16.667-2.792 16.667-6.25q.03-.26 0-.52"/><path stroke="#344054" d="M25 14.583c5.753 0 10.417-1.865 10.417-4.166S30.753 6.25 25 6.25s-10.417 1.865-10.417 4.167c0 2.3 4.664 4.166 10.417 4.166"/></g>`), g.Group(children),
	)
}

func BoatTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M33.854 16.667c2.38-.015 4.702-.74 6.667-2.084v0a2.084 2.084 0 0 1 3.229 1.792v2.5a8.333 8.333 0 0 1-8.208 8.208H14.458a8.333 8.333 0 0 1-8.208-8.208v-2.5a2.083 2.083 0 0 1 3.23-1.792v0a11.94 11.94 0 0 0 6.666 2.084z"/><path stroke="#344054" d="M33.792 10.417L16.667 34.979zM15.917 39.583l1.187-1.708a2.084 2.084 0 0 0-.437-2.896a2.085 2.085 0 0 0-2.917.5l-1.25 1.709z"/></g>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 6.25H16.667V25L25 22.917L33.333 25z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func BracketSquareTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 35.417h2.083a2.083 2.083 0 0 0 2.083-2.084v-6.25A2.083 2.083 0 0 1 35.417 25a2.083 2.083 0 0 1-2.084-2.083v-6.25a2.084 2.084 0 0 0-2.083-2.084h-2.083m-8.334 0H18.75a2.083 2.083 0 0 0-2.083 2.084v6.25A2.083 2.083 0 0 1 14.583 25a2.083 2.083 0 0 1 2.084 2.083v6.25a2.083 2.083 0 0 0 2.083 2.084h2.083"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func BracketTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 6.25h4.167A2.083 2.083 0 0 1 37.5 8.333v12.5A4.167 4.167 0 0 0 41.667 25a4.167 4.167 0 0 0-4.167 4.167v12.5a2.083 2.083 0 0 1-2.083 2.083H31.25"/><path stroke="#306CFE" d="M18.75 6.25h-4.167A2.083 2.083 0 0 0 12.5 8.333v12.5A4.167 4.167 0 0 1 8.333 25a4.167 4.167 0 0 1 4.167 4.167v12.5a2.083 2.083 0 0 0 2.083 2.083h4.167"/></g>`), g.Group(children),
	)
}

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.75 14.583C43.75 9.98 35.417 6.25 25 6.25S6.25 9.98 6.25 14.583c0 2.084 1.563 3.792 4.167 5.23v21.854A2.083 2.083 0 0 0 12.5 43.75h25a2.083 2.083 0 0 0 2.083-2.083V19.813c2.605-1.438 4.167-3.146 4.167-5.23"/>`), g.Group(children),
	)
}

func BreakfastTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 21.75c-4.792-1.062-6.25-7.167-10.896-4.312c-5.084 3.395-1.604 7.562.125 10.291s3.646 7.73 9.25 4.584c3.666-2.084 5.666-1.209 6.562-4.73c.896-3.52-2.395-5.229-5.041-5.833"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func BrickwallTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 29.167h16.666v-8.334H22.917zm-6.25 0H6.25V37.5h10.417zm0-16.667H6.25v8.333h10.417z"/><path stroke="#306CFE" d="M16.667 29.167h16.666V37.5H16.667zm-10.417 0h16.667v-8.334H6.25zm10.417-8.334h16.666V12.5H16.667zm27.083 8.334H33.333V37.5H43.75z"/></g>`), g.Group(children),
	)
}

func Brightness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m11.75 11.75l1.458 1.458M25 6.25v2.083zm13.25 5.5l-1.458 1.458zM43.75 25h-2.083zm-5.5 13.25l-1.458-1.458zM25 43.75v-2.083zm-13.25-5.5l1.458-1.458zM6.25 25h2.083z"/><path stroke="#306CFE" d="M25 33.333a8.333 8.333 0 1 0 0-16.666a8.333 8.333 0 0 0 0 16.666"/></g>`), g.Group(children),
	)
}

func Browse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 25H6.25zm-10.417 0c0-10.417-3.729-18.75-8.333-18.75S16.667 14.583 16.667 25S20.396 43.75 25 43.75s8.333-8.333 8.333-18.75"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 37.5A10.42 10.42 0 0 1 39 43.75M16.667 37.5A10.42 10.42 0 0 0 11 43.75m24.875-37.5a5.88 5.88 0 0 1-3.792 4.27l-1.854.647m-10.459 0l-1.853-.646a5.87 5.87 0 0 1-3.792-4.271m29.625 8.542l-1.104 1.875a11.3 11.3 0 0 1-7.23 5.146v0m.001 6.312a10.54 10.54 0 0 1 7.604 4.813l.729 1.145m-29.167-12.27a11.3 11.3 0 0 1-7.229-5.146L6.25 14.792m0 19.291l.73-1.145a10.54 10.54 0 0 1 7.603-4.813v0"/><path stroke="#306CFE" d="M35.417 31.25V20.833a2.083 2.083 0 0 0-2.084-2.083H16.667a2.083 2.083 0 0 0-2.084 2.083V31.25A10.417 10.417 0 0 0 25 41.667v0A10.417 10.417 0 0 0 35.417 31.25m-4.167-12.5h-12.5v-4.167A6.25 6.25 0 0 1 25 8.333v0a6.25 6.25 0 0 1 6.25 6.25z"/></g>`), g.Group(children),
	)
}

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M10.23 35.417H8.332a2.083 2.083 0 0 1-2.083-2.084V12.5a2.084 2.084 0 0 1 2.083-2.083h29.75a2.08 2.08 0 0 1 1.98 1.416l3.687 10.771l.104.646v10.083a2.083 2.083 0 0 1-2.083 2.084h-1.896"/><path stroke="#306CFE" d="M6.25 22.917h37.5zm25 12.5H18.917zm-12.5-12.5h12.5v-12.5h-12.5z"/><path stroke="#344054" d="M14.583 31.25a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334m20.834 0a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func Cabinet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.5 39.583v4.167M22.917 14.583h4.166zm0 16.667h4.166zM12.5 39.583v4.167z"/><path stroke="#306CFE" d="M10.417 22.917h29.166V8.333A2.083 2.083 0 0 0 37.5 6.25h-25a2.083 2.083 0 0 0-2.083 2.083zM12.5 39.583h25a2.083 2.083 0 0 0 2.083-2.083V22.917H10.417V37.5a2.083 2.083 0 0 0 2.083 2.083"/></g>`), g.Group(children),
	)
}

func CabinetFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 16.667v-2.084h8.334v2.084m-8.334 18.75v-2.084h8.334v2.084"/><path stroke="#306CFE" d="M8.333 25h33.334m0-16.667v33.334a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083V8.333a2.083 2.083 0 0 1 2.084-2.083h29.166a2.083 2.083 0 0 1 2.084 2.083"/></g>`), g.Group(children),
	)
}

func CabinetFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 25v-4.167m-6.25 18.75v4.167zm29.166 0v4.167zm-6.25-18.75V25z"/><path stroke="#306CFE" d="M10.417 39.583h29.166a2.083 2.083 0 0 0 2.084-2.083V8.333a2.083 2.083 0 0 0-2.084-2.083H10.417a2.083 2.083 0 0 0-2.084 2.083V37.5a2.083 2.083 0 0 0 2.084 2.083m0 0H25V6.25H10.417a2.083 2.083 0 0 0-2.084 2.083V37.5a2.083 2.083 0 0 0 2.084 2.083"/></g>`), g.Group(children),
	)
}

func CablewayTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 43.75V25a2.083 2.083 0 0 1 2.083-2.083h8.334A2.083 2.083 0 0 1 31.25 25v18.75"/><path stroke="#306CFE" d="M25 6.25v8.333M41.667 29.73l-3.563 12.5a2.084 2.084 0 0 1-2.083 1.521H13.979a2.08 2.08 0 0 1-2.083-1.52l-3.563-12.5a2.1 2.1 0 0 1 0-1.126l3.563-12.5a2.08 2.08 0 0 1 2.083-1.52h22.042a2.08 2.08 0 0 1 2.083 1.52l3.563 12.5c.103.368.103.757 0 1.125M14.583 6.25h20.834z"/></g>`), g.Group(children),
	)
}

func Call(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583A10.417 10.417 0 0 1 35.417 25m-12.5-18.625A17 17 0 0 1 25 6.25A18.75 18.75 0 0 1 43.75 25q0 1.045-.125 2.083"/><path stroke="#306CFE" d="M31.25 43.75a27.5 27.5 0 0 1-9.042-3.98L26 34.334a4.167 4.167 0 0 1 5.583-1.062c.688.437 1.396.833 2.084 1.229a4.167 4.167 0 0 1 1.75 5.875l-1.292 1.98a2.44 2.44 0 0 1-2.875 1.395M7.646 15.98l2.083-1.293a4.167 4.167 0 0 1 5.938 1.521c.396.75.791 1.459 1.229 2.084a4.167 4.167 0 0 1-1.063 5.583l-5.541 3.854A27.5 27.5 0 0 1 6.25 18.75a2.44 2.44 0 0 1 1.396-2.77"/><path stroke="#306CFE" d="M10.292 27.73a37 37 0 0 0 5.375 6.603a37 37 0 0 0 6.604 5.375"/></g>`), g.Group(children),
	)
}

func CallInTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25c9.646.563 15.438 8.333 15.438 18.75"/><path stroke="#344054" d="M35.146 21.417L40.437 25l3.313-5.708"/><path stroke="#306CFE" d="M31.25 43.75a27.5 27.5 0 0 1-9.042-3.98L26 34.334a4.167 4.167 0 0 1 5.583-1.062c.688.437 1.396.833 2.084 1.229a4.167 4.167 0 0 1 1.75 5.875l-1.292 1.98a2.44 2.44 0 0 1-2.875 1.395M7.646 15.98l2.083-1.293a4.167 4.167 0 0 1 5.938 1.521c.396.75.791 1.459 1.229 2.084a4.167 4.167 0 0 1-1.063 5.583l-5.541 3.854A27.5 27.5 0 0 1 6.25 18.75a2.44 2.44 0 0 1 1.396-2.77"/><path stroke="#306CFE" d="M10.292 27.73a37 37 0 0 0 5.375 6.603a37 37 0 0 0 6.604 5.375"/></g>`), g.Group(children),
	)
}

func Camping(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.27 33.854l-6.603 9.896h16.666l-6.604-9.896a2.083 2.083 0 0 0-3.458 0"/><path stroke="#306CFE" d="m20.833 6.25l20.98 34.333a2.085 2.085 0 0 1-1.771 3.167H9.958a2.083 2.083 0 0 1-1.77-3.167L29.167 6.25"/></g>`), g.Group(children),
	)
}

func CampingChair(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 27.667V8.333a2.083 2.083 0 0 1 2.083-2.083h20.834A2.083 2.083 0 0 1 37.5 8.333v19.334m-25 0L35.417 43.75m2-16.02L14.583 43.75"/><path stroke="#344054" d="M8.333 27.083a106.5 106.5 0 0 0 33.334 0"/></g>`), g.Group(children),
	)
}

func CampingGas(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 31.25v-4.167m-8.334 0v4.167"/><path stroke="#306CFE" d="M14.583 31.25h20.834a2.083 2.083 0 0 1 2.083 2.083V43.75h-25V33.333a2.083 2.083 0 0 1 2.083-2.083m6.25-16.667A3.895 3.895 0 0 0 25 18.75a3.957 3.957 0 0 0 4.167-4.167c0-2.791-2.084-3.125-2.084-8.333c-4.166 2.083-6.25 5.396-6.25 8.333"/><path stroke="#344054" d="M8.333 43.75h33.334"/></g>`), g.Group(children),
	)
}

func CannedFood(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 35.417h12.5m5.146-8.334a13.542 13.542 0 1 0-22.792 0z"/><path stroke="#306CFE" d="M39.583 27.083H10.417c-1.15 0-2.084.933-2.084 2.084v12.5c0 1.15.933 2.083 2.084 2.083h29.166a2.084 2.084 0 0 0 2.084-2.083v-12.5a2.084 2.084 0 0 0-2.084-2.084"/></g>`), g.Group(children),
	)
}

func CannedFoodTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M36.396 27.083a13.542 13.542 0 1 0-22.792 0z"/><path stroke="#306CFE" d="M39.583 27.083H10.417a2.084 2.084 0 0 0-2.084 2.084v12.5c0 1.15.933 2.083 2.084 2.083h29.166c1.15 0 2.084-.933 2.084-2.083v-12.5a2.084 2.084 0 0 0-2.084-2.084"/></g>`), g.Group(children),
	)
}

func CarAllert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M29.063 23.958h.208"/><path stroke="#306CFE" stroke-width="2" d="M39.583 37.5h2.084a2.083 2.083 0 0 0 2.083-2.083v-6.709a2.08 2.08 0 0 0-1.583-2.083l-2.584-.583m-18.75-9.375h-4.958A2.08 2.08 0 0 0 14 17.812L10.417 25l-3.021 1.5a2.08 2.08 0 0 0-1.146 1.875v7.042A2.083 2.083 0 0 0 8.333 37.5h2.084m8.333 0h12.5"/><path stroke="#344054" stroke-width="2" d="M29.167 8.333v6.25M18.75 37.5a4.167 4.167 0 1 0-8.334 0a4.167 4.167 0 0 0 8.334 0m20.833 0a4.166 4.166 0 1 0-8.332 0a4.166 4.166 0 0 0 8.332 0"/></g>`), g.Group(children),
	)
}

func CarLifter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 43.75l20.834-14.583l-20.834-14.584"/><path stroke="#344054" d="M35.417 14.583L14.583 29.167L35.417 43.75"/><path stroke="#306CFE" d="M6.25 43.75h37.5m-2.083-29.167H8.333A2.083 2.083 0 0 1 6.25 12.5V8.333A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083V12.5a2.083 2.083 0 0 1-2.083 2.083"/></g>`), g.Group(children),
	)
}

func CarShipping(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M31.25 12.5v22.917H20.833a4.167 4.167 0 0 0-8.333 0H8.333a2.083 2.083 0 0 1-2.083-2.084V12.5a2.083 2.083 0 0 1 2.083-2.083h20.834A2.083 2.083 0 0 1 31.25 12.5m12.146 9.896l-3.188-4.792a2.08 2.08 0 0 0-1.75-.937H31.25v18.75h2.083a4.167 4.167 0 0 1 8.334 0h2.083V23.542a2.1 2.1 0 0 0-.354-1.146"/><path stroke="#344054" d="M20.833 35.417a4.166 4.166 0 1 1-8.331 0a4.166 4.166 0 0 1 8.331 0M37.5 31.25a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M38.542 33.333h6.25M36.27 18.208a8.33 8.33 0 0 0-7.896-5.708H7.292a2.083 2.083 0 0 0-2.084 2.083V31.25a2.083 2.083 0 0 0 2.084 2.083h4.166a4.167 4.167 0 0 1 8.334 0h18.75V25z"/><path stroke="#344054" d="M28.125 20.833h-4.167m-4.166 12.5a4.167 4.167 0 1 1-8.335 0a4.167 4.167 0 0 1 8.335 0"/></g>`), g.Group(children),
	)
}

func CargoShipTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 31.25h-4.166M12.5 10.417v4.166z"/><path stroke="#306CFE" d="M6.25 22.917h35.417a2.083 2.083 0 0 1 1.958 2.812l-4.687 12.5a2.08 2.08 0 0 1-1.959 1.354h-25A2.08 2.08 0 0 1 9.896 38zm16.667 0l-1.688-6.75a2.08 2.08 0 0 0-2.083-1.584H8.333a2.083 2.083 0 0 0-2.083 2.084v6.25z"/></g>`), g.Group(children),
	)
}

func CarpetThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 33.333H31.25v-12.5H14.583z"/><path stroke="#306CFE" d="M10.417 41.667H37.5a2.083 2.083 0 0 0 2.083-2.084v-25A2.083 2.083 0 0 0 37.5 12.5H14.583"/><path stroke="#306CFE" d="M43.75 35.417h-4.167m-25-22.917v25a4.167 4.167 0 1 1-8.333 0v-25a4.167 4.167 0 1 1 8.333 0m-4.166 20.833a4.166 4.166 0 1 0 0 8.333a4.166 4.166 0 0 0 0-8.333M43.75 18.75h-4.167zm0 8.333h-4.167z"/></g>`), g.Group(children),
	)
}

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M35.417 42.708h.208m-12.708 0h.208z"/><path stroke="#306CFE" stroke-width="2" d="M6.25 6.25h4.458a2.08 2.08 0 0 1 2.084 1.77l1 6.563l2.875 18.75l22.916-2.083l4.167-16.667H13.792"/></g>`), g.Group(children),
	)
}

func CartAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M35.417 42.708h.208m-12.708 0h.208z"/><path stroke="#344054" stroke-width="2" d="M35.417 16.667h-12.5m6.25-6.25v12.5z"/><path stroke="#306CFE" stroke-width="2" d="M6.25 6.25h4.583a2.084 2.084 0 0 1 2.084 1.625l5.458 23.833a2.08 2.08 0 0 0 2.083 1.625H38a2.084 2.084 0 0 0 2.083-1.583l3.667-15.083"/></g>`), g.Group(children),
	)
}

func CartAltOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M35.417 42.708h.208m-12.708 0h.208z"/><path stroke="#306CFE" stroke-width="2" d="M6.25 6.25h4.833a2.08 2.08 0 0 1 1.938 1.313l7.812 19.52l-2.666 5.313a2.084 2.084 0 0 0 1.875 3.02h19.541"/><path stroke="#306CFE" stroke-width="2" d="M15 12.5h28.75l-5.833 14.583H20.833"/></g>`), g.Group(children),
	)
}

func CartRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M35.417 42.708h.208m-12.708 0h.208z"/><path stroke="#344054" stroke-width="2" d="M35.417 16.667h-12.5"/><path stroke="#306CFE" stroke-width="2" d="M6.25 6.25h4.583a2.084 2.084 0 0 1 2.084 1.625l5.458 23.833a2.08 2.08 0 0 0 2.083 1.625H38a2.084 2.084 0 0 0 2.083-1.583l3.667-15.083"/></g>`), g.Group(children),
	)
}

func CashierMachineLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 22.917v-8.334m8.333-6.25V12.5A2.083 2.083 0 0 1 25 14.583H12.5a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 12.5 6.25H25a2.083 2.083 0 0 1 2.083 2.083"/><path stroke="#306CFE" d="M6.25 35.417v6.25a2.083 2.083 0 0 0 2.083 2.083h33.334a2.083 2.083 0 0 0 2.083-2.083v-6.25l-3.687-11.084a2.08 2.08 0 0 0-1.98-1.416H11.917a2.08 2.08 0 0 0-2.084 1.416zm0 0v6.25a2.083 2.083 0 0 0 2.083 2.083h33.334a2.083 2.083 0 0 0 2.083-2.083v-6.25z"/></g>`), g.Group(children),
	)
}

func Cassette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 18.75H14.583m18.75 12.5H16.667l-2.084 8.333h20.834z"/><path stroke="#306CFE" d="M41.667 10.417H8.333c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M27.5 27.083h-5a2.083 2.083 0 0 1-2.083-1.875L18.75 8.542a2.084 2.084 0 0 1 2.083-2.292h8.334a2.083 2.083 0 0 1 2.083 2.292l-1.667 16.666a2.083 2.083 0 0 1-2.083 1.875"/><path stroke="#344054" d="M25 43.75a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/></g>`), g.Group(children),
	)
}

func CautionSignCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 34.375h.208"/><path stroke="#306CFE" stroke-width="2" d="M43.75 25c0-10.355-8.395-18.75-18.75-18.75S6.25 14.645 6.25 25S14.645 43.75 25 43.75S43.75 35.355 43.75 25"/><path stroke="#344054" stroke-width="2" d="M25 25V14.583"/></g>`), g.Group(children),
	)
}

func CautionSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 34.375h.208"/><path stroke="#306CFE" stroke-width="2" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/><path stroke="#344054" stroke-width="2" d="M25 25V14.583"/></g>`), g.Group(children),
	)
}

func Cd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 29.167a4.167 4.167 0 1 0 0-8.334a4.167 4.167 0 0 0 0 8.334"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func ChairFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 6.25v20.833m0 8.334v8.333zm20.834 0v8.333zm0-8.334V6.25z"/><path stroke="#306CFE" d="M35.417 18.75H14.583v-8.333h20.834zM37.5 35.417h-25v-6.25a2.083 2.083 0 0 1 2.083-2.084h20.834a2.083 2.083 0 0 1 2.083 2.084z"/></g>`), g.Group(children),
	)
}

func ChairTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 35.417v8.333m-20.834-8.333v8.333z"/><path stroke="#306CFE" d="M35.417 27.083H14.583V8.333a2.083 2.083 0 0 1 2.084-2.083h16.666a2.083 2.083 0 0 1 2.084 2.083zm2.083 8.334h-25v-6.25a2.083 2.083 0 0 1 2.083-2.084h20.834a2.083 2.083 0 0 1 2.083 2.084z"/></g>`), g.Group(children),
	)
}

func ChaiseLongue(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 27.083a8.333 8.333 0 0 0-8.333-8.333H25v-1.042a7.292 7.292 0 0 0-14.583 0v.563m29.166 21.312v-4.166m-29.166 4.166v-4.166z"/><path stroke="#306CFE" d="M14.583 27.083v-3.937a4.31 4.31 0 0 0-3.458-4.396a4.168 4.168 0 0 0-4.875 4.167v12.5h37.5v-6.25a2.083 2.083 0 0 0-2.083-2.084z"/></g>`), g.Group(children),
	)
}

func ChartColum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M28.125 6.25h-6.25v37.5h6.25z"/><path stroke="#306CFE" d="M12.5 43.75H6.25V25h6.25zm31.25-29.167H37.5V43.75h6.25z"/></g>`), g.Group(children),
	)
}

func ChartColumnGrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 25c8.333 0 27.083-4.167 35.417-18.75"/><path stroke="#344054" d="M33.333 6.25h8.334l2.083 8.333"/><path stroke="#306CFE" d="M12.5 43.75H6.25v-8.333h6.25zm16.667-12.5h-6.25v12.5h6.25zm16.666-8.333h-6.25V43.75h6.25z"/></g>`), g.Group(children),
	)
}

func ChartColumnLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 6.25c7.208 12.02 22.75 16.104 31.25 16.667"/><path stroke="#344054" d="m41.667 16.667l2.083 6.25l-6.25 4.166"/><path stroke="#306CFE" d="M39.583 35.417h6.25v8.333h-6.25zM22.917 43.75h6.25v-12.5h-6.25zm-16.667 0h6.25V22.917H6.25z"/></g>`), g.Group(children),
	)
}

func ChartLineAltOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M14.583 33.333L25 22.917l8.333 6.25l7.396-14.813"/><path stroke="#344054" stroke-width="3" d="M41.563 12.5h.208"/><path stroke="#306CFE" stroke-width="2" d="M6.25 6.25v37.5h37.5"/></g>`), g.Group(children),
	)
}

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M32.354 38.792a14.98 14.98 0 1 1-11.52-25.021"/><path stroke="#306CFE" d="M20.833 29.167V6.25A23.04 23.04 0 0 1 43.75 29.167a22.92 22.92 0 0 1-5.333 14.687z"/></g>`), g.Group(children),
	)
}

func Chassis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 12.5v25m-8.333-25h16.666zm0 25h16.666z"/><path stroke="#306CFE" d="M37.5 18.75a4.167 4.167 0 0 1-4.167-4.167v-4.166A4.167 4.167 0 0 1 37.5 6.25v0a4.167 4.167 0 0 1 4.167 4.167v4.166A4.167 4.167 0 0 1 37.5 18.75m-20.833-4.167v-4.166a4.167 4.167 0 1 0-8.334 0v4.166a4.167 4.167 0 1 0 8.334 0m25 25v-4.166a4.167 4.167 0 1 0-8.334 0v4.166a4.167 4.167 0 1 0 8.334 0m-25 0v-4.166a4.167 4.167 0 1 0-8.334 0v4.166a4.167 4.167 0 1 0 8.334 0"/></g>`), g.Group(children),
	)
}

func ChatAlertLeftThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 28.125h-.208"/><path stroke="#344054" stroke-width="2" d="M25 18.75v-4.167"/><path stroke="#306CFE" stroke-width="2" d="M41.667 6.25H8.333A2.083 2.083 0 0 0 6.25 8.333v27.084A2.083 2.083 0 0 0 8.333 37.5h8.334v6.25l10.416-6.25h14.584a2.083 2.083 0 0 0 2.083-2.083V8.333a2.083 2.083 0 0 0-2.083-2.083"/></g>`), g.Group(children),
	)
}

func ChatAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.75 8.333v27.084a2.083 2.083 0 0 1-2.083 2.083H31.25L25 43.75l-6.25-6.25H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083"/>`), g.Group(children),
	)
}

func ChatAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M33.333 22.917h.209m-8.542 0h.208zm-8.542 0h.209z"/><path stroke="#306CFE" stroke-width="2" d="m39.188 33.813l2.479 9.937l-10.313-5.167a20.5 20.5 0 0 1-6.354 1c-10.417 0-18.75-7.458-18.75-16.666S14.583 6.25 25 6.25s18.75 7.458 18.75 16.667a15.6 15.6 0 0 1-4.562 10.895"/></g>`), g.Group(children),
	)
}

func ChatFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 41.667h-6.25a2.083 2.083 0 0 1-2.083-2.084V16.667M22.917 25h12.5m-12.5-8.333h12.5z"/><path stroke="#306CFE" d="M41.667 8.333h-25a2.083 2.083 0 0 0-2.084 2.084V31.25a2.083 2.083 0 0 0 2.084 2.083h6.25v6.25l10.416-6.25h8.334a2.083 2.083 0 0 0 2.083-2.083V10.417a2.083 2.083 0 0 0-2.083-2.084"/></g>`), g.Group(children),
	)
}

func ChatLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 41.667h-6.25a2.083 2.083 0 0 1-2.083-2.084V16.667"/><path stroke="#306CFE" d="M41.667 8.333h-25a2.083 2.083 0 0 0-2.084 2.084V31.25a2.083 2.083 0 0 0 2.084 2.083h6.25v6.25l10.416-6.25h8.334a2.083 2.083 0 0 0 2.083-2.083V10.417a2.083 2.083 0 0 0-2.083-2.084"/></g>`), g.Group(children),
	)
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m16.667 25l6.25 6.25l10.416-10.417"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func CheckDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m27.542 34.375l16.208-18.75"/><path stroke="#306CFE" d="m6.25 24.75l8.333 9.625l16.209-18.75"/></g>`), g.Group(children),
	)
}

func CheckMarkCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m16.667 23.958l6.25 6.25l10.416-10.416"/><path stroke="#306CFE" d="M43.75 25c0-10.355-8.395-18.75-18.75-18.75S6.25 14.645 6.25 25S14.645 43.75 25 43.75S43.75 35.355 43.75 25"/></g>`), g.Group(children),
	)
}

func CheckMarkSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 19.792L22.917 30.208l-6.25-6.25"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func CheckMarkSquareTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 10.417L25 29.167l-8.333-8.334"/><path stroke="#306CFE" d="M43.75 22.917v18.75a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h25"/></g>`), g.Group(children),
	)
}

func Chemistry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 43.75H12.5a2.083 2.083 0 0 1-2.083-2.083V6.25M6.25 14.583h20.833"/><path stroke="#344054" d="M25 8.333h12.5m-2.083 20.834V8.333h-8.334v20.834a4.167 4.167 0 1 0 8.334 0"/></g>`), g.Group(children),
	)
}

func ChemistryFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 43.75H12.5a2.083 2.083 0 0 1-2.083-2.083V6.25M6.25 14.583h20.833"/><path stroke="#344054" d="M25 8.333h12.5m3.917 24.021l-6-11.146V8.333h-8.334v12.875l-6 11.146a2.083 2.083 0 0 0 1.834 3.063h16.666a2.082 2.082 0 0 0 1.834-3.063"/></g>`), g.Group(children),
	)
}

func Chip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 33.333H8.333M25 12.5V6.25zm8.333 0V8.333zm-16.666 0V8.333zM37.5 25h6.25zm0 8.333h4.167zm0-16.666h4.167zM25 37.5v6.25zm-8.333 0v4.167zm16.666 0v4.167zM12.5 25H6.25zm0-8.333H8.333z"/><path stroke="#306CFE" d="M35.417 12.5H14.583c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.933 2.083 2.083 2.083h20.834c1.15 0 2.083-.933 2.083-2.083V14.583c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func ChronometerWatchThreeSecond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 6.25h12.5M25 12.5V6.25zm0 8.333v8.334h4.167"/><path stroke="#306CFE" d="M25 43.75c8.63 0 15.625-6.996 15.625-15.625c0-8.63-6.996-15.625-15.625-15.625c-8.63 0-15.625 6.996-15.625 15.625c0 8.63 6.996 15.625 15.625 15.625"/></g>`), g.Group(children),
	)
}

func ClipboardAddTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 27.083V8.333a2.083 2.083 0 0 0-2.083-2.083h-25a2.083 2.083 0 0 0-2.084 2.083v33.334a2.083 2.083 0 0 0 2.084 2.083H25"/><path stroke="#344054" d="M37.5 35.417v8.333M16.667 6.25h12.5v6.25a2.083 2.083 0 0 1-2.084 2.083H18.75a2.083 2.083 0 0 1-2.083-2.083zm0 29.167H25zm0-8.334h12.5zm16.666 12.5h8.334z"/></g>`), g.Group(children),
	)
}

func ClipboardChecklistTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 6.25h-25c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.932 2.083 2.083 2.083h25c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.932-2.083-2.083-2.083"/><path stroke="#344054" d="m18.75 29.167l4.167 4.166L31.25 25m-12.5-12.5a2.083 2.083 0 0 0 2.083 2.083h8.334A2.083 2.083 0 0 0 31.25 12.5V6.25h-12.5z"/></g>`), g.Group(children),
	)
}

func ClipboardDeleteTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 31.25V8.333a2.083 2.083 0 0 0-2.083-2.083h-25a2.083 2.083 0 0 0-2.084 2.083v33.334a2.083 2.083 0 0 0 2.084 2.083H25"/><path stroke="#344054" d="M33.333 39.583h8.334m-25-33.333h12.5v6.25a2.083 2.083 0 0 1-2.084 2.083H18.75a2.083 2.083 0 0 1-2.083-2.083zm0 29.167H25zm0-8.334h12.5z"/></g>`), g.Group(children),
	)
}

func ClipboardEditLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M14.583 33.333v8.334a2.083 2.083 0 0 0 2.084 2.083h25a2.083 2.083 0 0 0 2.083-2.083V8.333a2.083 2.083 0 0 0-2.083-2.083h-25a2.083 2.083 0 0 0-2.084 2.083v12.5"/><path stroke="#344054" d="M35.417 12.5a2.083 2.083 0 0 1-2.084 2.083H25a2.083 2.083 0 0 1-2.083-2.083V6.25h12.5zM6.854 22.27l2.917-2.916a2.083 2.083 0 0 1 2.916 0l10.23 10.23v5.833h-5.834L6.854 25.187a2.083 2.083 0 0 1 0-2.916"/></g>`), g.Group(children),
	)
}

func ClipboardEditThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 33.333v8.334a2.083 2.083 0 0 1-2.084 2.083h-25a2.083 2.083 0 0 1-2.083-2.083V14.583A2.083 2.083 0 0 1 8.333 12.5h6.25m12.5 0h6.25a2.083 2.083 0 0 1 2.084 2.083v6.25"/><path stroke="#344054" d="M25 10.417a4.167 4.167 0 1 0-8.333 0v0a2.083 2.083 0 0 0-2.084 2.083v4.167h12.5V12.5A2.083 2.083 0 0 0 25 10.417M43.146 22.27l-2.917-2.917a2.083 2.083 0 0 0-2.916 0l-10.23 10.23v5.833h5.834l10.229-10.23a2.083 2.083 0 0 0 0-2.916"/></g>`), g.Group(children),
	)
}

func ClockAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 16.667V25l4.167 4.167m4.166 10.416L37.5 43.75m-25-37.5a6.25 6.25 0 0 0-2.687 11.875a16.67 16.67 0 0 1 8.333-8.333A6.25 6.25 0 0 0 12.5 6.25m25 0a6.25 6.25 0 0 0-5.625 3.563a16.67 16.67 0 0 1 8.333 8.333A6.25 6.25 0 0 0 37.5 6.25M16.667 39.583L12.5 43.75z"/><path stroke="#306CFE" d="M41.667 25a16.667 16.667 0 0 1-33.334 0a16.5 16.5 0 0 1 1.48-6.875a16.67 16.67 0 0 1 8.333-8.333a16.67 16.67 0 0 1 13.75 0a16.67 16.67 0 0 1 8.333 8.333A16.5 16.5 0 0 1 41.667 25"/></g>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 31.48a13.73 13.73 0 1 1 20.833-15.084a9.6 9.6 0 0 1 1.73-.167A8.73 8.73 0 0 1 37.5 33.333"/><path stroke="#306CFE" d="M33.056 31.611h-.558a.945.945 0 0 1-.887-.633a.94.94 0 0 1 .188-1.076l.388-.388a.945.945 0 0 0 0-1.34l-1.341-1.351a.946.946 0 0 0-1.342 0l-.387.387a.94.94 0 0 1-1.076.189a.945.945 0 0 1-.652-.897v-.558a.945.945 0 0 0-.945-.944h-1.888a.945.945 0 0 0-.945.944v.558a.945.945 0 0 1-.633.887v0a.94.94 0 0 1-1.076-.188l-.388-.388a.945.945 0 0 0-1.34 0l-1.351 1.341a.946.946 0 0 0 0 1.342l.387.387a.94.94 0 0 1 .189 1.076a.945.945 0 0 1-.888.633h-.567a.945.945 0 0 0-.944.945v1.889a.945.945 0 0 0 .944.944h.558a.94.94 0 0 1 .887.633v0a.94.94 0 0 1-.188 1.076l-.388.388a.945.945 0 0 0 0 1.34l1.332 1.332a.945.945 0 0 0 1.341 0l.387-.387a.945.945 0 0 1 1.077-.189a.945.945 0 0 1 .633.888v.605a.945.945 0 0 0 .944.944h1.89a.944.944 0 0 0 .944-.944v-.558a.94.94 0 0 1 .632-.887a.94.94 0 0 1 1.077.188l.387.388a.946.946 0 0 0 1.341 0l1.332-1.332a.945.945 0 0 0 0-1.341l-.387-.387a.95.95 0 0 1-.19-1.077v0a.94.94 0 0 1 .889-.633h.614a.945.945 0 0 0 .944-.944v-1.917a.945.945 0 0 0-.944-.945"/><path stroke="#344054" d="M25.5 36a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/></g>`), g.Group(children),
	)
}

func CoffeeMachine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 14.583v4.167m4.167 12.5v-4.167h-8.333v4.167a4.167 4.167 0 1 0 8.333 0"/><path stroke="#306CFE" d="M41.667 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083v-6.25h37.5v6.25a2.083 2.083 0 0 1-2.083 2.083m-2.084-37.5H8.333A2.083 2.083 0 0 0 6.25 8.333v27.084h14.583V14.583H37.5a2.083 2.083 0 0 0 2.083-2.083z"/></g>`), g.Group(children),
	)
}

func CompactDiskDisable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/><path stroke="#344054" d="m8.333 8.333l33.334 33.334M25 20.833a4.166 4.166 0 1 0 0 8.333a4.166 4.166 0 0 0 0-8.333"/></g>`), g.Group(children),
	)
}

func CompactDiskThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M35.417 25A10.417 10.417 0 0 0 25 14.583M14.583 25A10.417 10.417 0 0 0 25 35.417"/><path stroke="#344054" stroke-width="3" d="M24.896 25h.208"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func CompactDiskTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M38.25 11.75a18.2 18.2 0 0 0-7.312-4.52l-5.73 13.603a4.167 4.167 0 0 1 3.959 3.938l13.625-5.73a18.2 18.2 0 0 0-4.542-7.29"/><path stroke="#344054" d="M25 29.167a4.167 4.167 0 1 0 0-8.334a4.167 4.167 0 0 0 0 8.334"/><path stroke="#344054" d="m19.063 42.77l5.729-13.603a4.166 4.166 0 0 1-3.959-3.938L7.23 30.938A18.75 18.75 0 0 0 19.063 42.77"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func ConeGeometric(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 43.75c8.054 0 14.583-2.798 14.583-6.25S33.054 31.25 25 31.25s-14.583 2.798-14.583 6.25S16.946 43.75 25 43.75"/><path stroke="#306CFE" d="M39.292 36.27L25 6.25L10.708 36.27a2.9 2.9 0 0 0-.291 1.23c0 3.458 6.52 6.25 14.583 6.25s14.583-2.792 14.583-6.25a2.9 2.9 0 0 0-.291-1.23"/></g>`), g.Group(children),
	)
}

func ContractRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 27.083v14.584a2.083 2.083 0 0 1-2.084 2.083h-25a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h25a2.083 2.083 0 0 1 2.084 2.083v6.25"/><path stroke="#344054" d="M14.583 35.417h6.25M43.146 16.02l-2.917-2.917a2.083 2.083 0 0 0-2.916 0l-10.23 10.23v5.833h5.834l10.229-10.23a2.083 2.083 0 0 0 0-2.916"/></g>`), g.Group(children),
	)
}

func Contrass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 25A8.333 8.333 0 0 1 25 33.333V16.667A8.334 8.334 0 0 1 33.333 25"/><path stroke="#306CFE" d="m44.146 23.521l-4.563-4.562V12.5a2.083 2.083 0 0 0-2.083-2.083h-6.458l-4.563-4.563a2.084 2.084 0 0 0-2.958 0l-4.563 4.563H12.5a2.083 2.083 0 0 0-2.083 2.083v6.459L5.854 23.52a2.08 2.08 0 0 0 0 2.958l4.563 4.563V37.5a2.084 2.084 0 0 0 2.083 2.083h6.458l4.563 4.563a2.084 2.084 0 0 0 2.958 0l4.563-4.563H37.5a2.083 2.083 0 0 0 2.083-2.083v-6.458l4.563-4.563a2.084 2.084 0 0 0 0-2.958"/></g>`), g.Group(children),
	)
}

func ContrassAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583v20.834a10.417 10.417 0 1 0 0-20.834"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func ContrassAltSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583H14.583v20.834H25z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func ContrassAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583H14.583v20.834H25z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Conversation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.188 17.104a13.06 13.06 0 0 1 6.562 11.021a12.13 12.13 0 0 1-2.396 7.292l2.396 8.333l-8.333-3.604a17.7 17.7 0 0 1-7.292 1.52a15.56 15.56 0 0 1-14.812-9.27"/><path stroke="#306CFE" d="M37.5 19.792c0 7.479-7 13.541-15.625 13.541a17.7 17.7 0 0 1-7.292-1.52l-1.354.583l-6.979 3.02l2.396-8.333a12.13 12.13 0 0 1-2.396-7.291c0-7.48 7-13.542 15.625-13.542c7.563 0 13.875 4.667 15.313 10.854c.21.88.314 1.783.312 2.688"/></g>`), g.Group(children),
	)
}

func Couldron(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m39.583 39.583l-3.979-7.958m-25.187 7.958l3.979-7.958"/><path stroke="#306CFE" d="M39.417 10.417H10.583a16.36 16.36 0 0 0-2.25 8.333a16.667 16.667 0 1 0 33.334 0c.01-2.93-.768-5.807-2.25-8.333"/><path stroke="#344054" d="M6.25 10.417h37.5"/></g>`), g.Group(children),
	)
}

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 36.458v3.125m0-29.166v3.125zm0 10.416v8.334z"/><path stroke="#306CFE" d="M43.75 12.5v25a2.083 2.083 0 0 1-2.083 2.083H8.333A2.083 2.083 0 0 1 6.25 37.5v-6.25a6.25 6.25 0 0 0 0-12.5V12.5a2.083 2.083 0 0 1 2.083-2.083h33.334A2.083 2.083 0 0 1 43.75 12.5"/></g>`), g.Group(children),
	)
}

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 33.333h.208"/><path stroke="#344054" stroke-width="2" d="M29.167 22.917h-8.334m8.334-8.334h-8.334z"/><path stroke="#306CFE" stroke-width="2" d="M35.417 6.25H14.583c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h20.834c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Cradle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 6.25h-4.166a12.5 12.5 0 0 0-12.5 12.5h16.666z"/><path stroke="#306CFE" d="m37.5 42.604l-2.187-11.812M12.5 42.604l2.188-11.791M41.667 18.75H8.333m31.25 4.167V18.75H10.417v4.167a8.333 8.333 0 0 0 8.333 8.333h12.5a8.334 8.334 0 0 0 8.333-8.333"/><path stroke="#344054" d="M41.667 41.667a67.9 67.9 0 0 1-33.334 0"/></g>`), g.Group(children),
	)
}

func CreateNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M41.667 8.333a4.354 4.354 0 0 0-6.146.25L21.188 22.917L18.75 31.25l8.333-2.437l14.334-14.23a4.356 4.356 0 0 0 .25-6.25"/><path stroke="#306CFE" d="M25 6.25H8.333A2.083 2.083 0 0 0 6.25 8.333v33.334a2.083 2.083 0 0 0 2.083 2.083h33.334a2.083 2.083 0 0 0 2.083-2.083V25"/></g>`), g.Group(children),
	)
}

func CreateNoteAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 20.833s7.354.855 11.333-3.708c4-4.625 3.25-12.958 3.25-12.958a17 17 0 0 0-5.27.375l-.98 3.791l-3.98-1.479q-.644.487-1.166 1.104c-4 4.542-3.187 12.875-3.187 12.875m0 0L25 25"/><path stroke="#306CFE" d="M25 6.25H8.333A2.083 2.083 0 0 0 6.25 8.333v33.334a2.083 2.083 0 0 0 2.083 2.083h33.334a2.083 2.083 0 0 0 2.083-2.083V25"/></g>`), g.Group(children),
	)
}

func CrossArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M6.25 25h37.5M25 43.75V6.25z"/><path stroke="#344054" d="M12.5 18.75L6.25 25l6.25 6.25M31.25 12.5L25 6.25l-6.25 6.25M37.5 31.25L43.75 25l-6.25-6.25M18.75 37.5L25 43.75l6.25-6.25"/></g>`), g.Group(children),
	)
}

func CrossCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m31.25 31.25l-12.5-12.5m0 12.5l12.5-12.5"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 12.5v25M25 31.25L8.333 37.5L25 43.75l16.667-6.25zM8.333 12.5v25zM25 18.75v25z"/><path stroke="#344054" d="M41.667 12.5L25 18.75L8.333 12.5L25 6.25z"/></g>`), g.Group(children),
	)
}

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M40.708 6.458L7.542 21.48a2.437 2.437 0 0 0 .791 4.563l12.98 2.625l2.645 13a2.437 2.437 0 0 0 4.563.812L43.54 9.292a2.084 2.084 0 0 0-2.833-2.834"/>`), g.Group(children),
	)
}

func CursorTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M24.583 19.542L8.938 35.167a2.084 2.084 0 0 0 0 2.937l2.958 2.959a2.084 2.084 0 0 0 2.937 0l15.625-15.646"/><path stroke="#306CFE" d="m30.458 25.417l2.75 5.104a2.083 2.083 0 0 0 3.855-.334L41.667 11A2.188 2.188 0 0 0 39 8.333l-19.187 4.604a2.083 2.083 0 0 0-.334 3.855l5.104 2.75"/></g>`), g.Group(children),
	)
}

func Curtains(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 22.917v8.333"/><path stroke="#306CFE" d="M8.333 22.917h33.334l-2.084-8.334H10.417zm0-8.334h33.334L39.583 6.25H10.417zm2.084 29.167h29.166V22.917H10.417zm33.333 0H6.25m37.5-37.5H6.25z"/></g>`), g.Group(children),
	)
}

func CurveArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 16.667c-14.791.583-20.833 12.77-20.833 25c0 0 8.083-11.855 20.833-12.5"/><path stroke="#306CFE" d="M27.083 29.167V37.5L42.98 24.52a2.083 2.083 0 0 0 0-3.207L27.083 8.333v8.334"/></g>`), g.Group(children),
	)
}

func CurveArrowRightNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m37.5 25l6.25-6.25l-6.25-6.25"/><path stroke="#344054" d="M43.75 18.75h-8.937a22.6 22.6 0 0 0-20.23 12.5"/><path stroke="#306CFE" d="M43.75 33.333v6.25a2.083 2.083 0 0 1-2.083 2.084H8.333a2.083 2.083 0 0 1-2.083-2.084V10.417a2.083 2.083 0 0 1 2.083-2.084H25"/></g>`), g.Group(children),
	)
}

func CurveArrowRightThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m31.25 33.333l11.688-8.791a2.083 2.083 0 0 0 0-3.25L31.25 12.5"/><path stroke="#306CFE" d="M18.75 33.333v-6.25c-8.896 0-12.5 10.417-12.5 10.417c0-8.77 2.23-18.75 12.5-18.75V12.5l14.083 10.417z"/></g>`), g.Group(children),
	)
}

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 43.75c6.904 0 12.5-1.865 12.5-4.167c0-2.3-5.596-4.166-12.5-4.166s-12.5 1.865-12.5 4.166S18.096 43.75 25 43.75"/><path stroke="#306CFE" d="M37.5 10.417v29.166c0 2.292-5.604 4.167-12.5 4.167s-12.5-1.875-12.5-4.167V10.417c0-2.292 5.604-4.167 12.5-4.167s12.5 1.875 12.5 4.167"/><path stroke="#344054" d="M25 14.583c6.904 0 12.5-1.865 12.5-4.166S31.904 6.25 25 6.25s-12.5 1.865-12.5 4.167c0 2.3 5.596 4.166 12.5 4.166"/></g>`), g.Group(children),
	)
}

func DashboardAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 18.75v22.917m-25-22.917h37.5z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func DashboardAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M33.333 18.75L25 31.25"/><path stroke="#344054" stroke-width="3" d="M25 31.208v.209"/><path stroke="#306CFE" stroke-width="2" d="M43.75 27.083a18.75 18.75 0 0 1-7 14.584h-23.5a18.75 18.75 0 1 1 30.5-14.584"/></g>`), g.Group(children),
	)
}

func DatabaseTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 18.75H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083v8.334a2.083 2.083 0 0 1-2.083 2.083m0 25H8.333a2.083 2.083 0 0 1-2.083-2.083v-8.334a2.083 2.083 0 0 1 2.083-2.083h33.334a2.083 2.083 0 0 1 2.083 2.083v8.334a2.083 2.083 0 0 1-2.083 2.083"/><path stroke="#344054" d="M41.667 18.75H8.333c-1.15 0-2.083.933-2.083 2.083v8.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083v-8.334c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Date(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083V20.833h37.5v20.834a2.083 2.083 0 0 1-2.083 2.083M43.75 12.5a2.083 2.083 0 0 0-2.083-2.083H8.333A2.083 2.083 0 0 0 6.25 12.5v8.333h37.5z"/><path stroke="#344054" d="M33.333 6.25v8.333M16.667 6.25v8.333"/></g>`), g.Group(children),
	)
}

func DateAltAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M15.708 39.583H8.333A2.083 2.083 0 0 1 6.25 37.5V10.417a2.083 2.083 0 0 1 2.083-2.084h33.334a2.083 2.083 0 0 1 2.083 2.084V37.5a2.083 2.083 0 0 1-2.083 2.083h-7.375"/><path stroke="#306CFE" d="M6.25 18.75h37.5m-18.75 0a12.5 12.5 0 1 0 0 25a12.5 12.5 0 0 0 0-25"/><path stroke="#344054" d="M20.833 31.25h8.334m4.166-25v6.25zm-16.666 0v6.25zM25 35.417v-8.334z"/></g>`), g.Group(children),
	)
}

func DateAltCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M15.708 39.583H8.333A2.083 2.083 0 0 1 6.25 37.5V10.417a2.083 2.083 0 0 1 2.083-2.084h33.334a2.083 2.083 0 0 1 2.083 2.084V37.5a2.083 2.083 0 0 1-2.083 2.083h-7.375"/><path stroke="#306CFE" d="M6.25 18.75h37.5m-18.75 0a12.5 12.5 0 1 0 0 25a12.5 12.5 0 0 0 0-25"/><path stroke="#344054" d="M16.667 6.25v6.25m16.666-6.25v6.25zm-12.5 27.083l2.084 2.084l6.25-6.25"/></g>`), g.Group(children),
	)
}

func DateAltStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M15.708 39.583H8.333A2.083 2.083 0 0 1 6.25 37.5V10.417a2.083 2.083 0 0 1 2.083-2.084h33.334a2.083 2.083 0 0 1 2.083 2.084V37.5a2.083 2.083 0 0 1-2.083 2.083h-7.375"/><path stroke="#306CFE" d="M6.25 18.75h37.5m-18.75 0a12.5 12.5 0 1 0 0 25a12.5 12.5 0 0 0 0-25"/><path stroke="#344054" d="M33.333 6.25v6.25zm-16.666 0v6.25zM25 34l2.708 1.417l-.52-3.021l2.083-2.084l-3.021-.437L25 27.083l-1.354 2.75l-3.021.438l2.083 2.083l-.416 3.063z"/></g>`), g.Group(children),
	)
}

func DebitPurchase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M30.313 30.208h-.209m4.375 0h-.208z"/><path stroke="#344054" stroke-width="2" d="M43.75 18.75H6.25"/><path stroke="#306CFE" stroke-width="2" d="M41.667 10.417H8.333c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func DecreaseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.833 25H16.167"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 14.583v-6.25A2.083 2.083 0 0 0 31.25 6.25h-12.5a2.083 2.083 0 0 0-2.083 2.083v6.25"/><path stroke="#306CFE" d="M8.333 14.583h33.334M37.5 41.667V14.583h-25v27.084a2.083 2.083 0 0 0 2.083 2.083h20.834a2.083 2.083 0 0 0 2.083-2.083"/></g>`), g.Group(children),
	)
}

func DeleteAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 22.917v12.5m8.333-20.834v-6.25A2.083 2.083 0 0 0 31.25 6.25h-12.5a2.083 2.083 0 0 0-2.083 2.083v6.25z"/><path stroke="#306CFE" d="M8.333 14.583h33.334zm27.23 27.23l1.937-27.23h-25l1.938 27.23a2.083 2.083 0 0 0 2.083 1.937h16.958a2.084 2.084 0 0 0 2.084-1.938"/></g>`), g.Group(children),
	)
}

func DeleteCollection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m16.417 16.417l8.833 8.833m-8.833 0l8.833-8.833"/><path stroke="#306CFE" d="M33.333 6.25h-25c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.084 2.083 2.084h25c1.15 0 2.084-.933 2.084-2.084v-25c0-1.15-.933-2.083-2.084-2.083"/><path stroke="#306CFE" d="M14.583 43.75h27.084a2.083 2.083 0 0 0 2.083-2.083v-31.25"/></g>`), g.Group(children),
	)
}

func DeliveryTruck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M19.104 35.417h7.98V12.5A2.083 2.083 0 0 0 25 10.417H10.417M6.25 27.083v6.25a2.083 2.083 0 0 0 2.083 2.084h1.813"/><path stroke="#306CFE" d="M30.98 35.417h-3.897V14.584h8.855a2.08 2.08 0 0 1 2.083 1.52l1.562 5.771l2.584.646a2.08 2.08 0 0 1 1.583 2.083v8.73a2.083 2.083 0 0 1-2.083 2.083h-1.855"/><path stroke="#344054" d="M6.25 18.75h12.5m-4.167 12.5a4.167 4.167 0 1 0 0 8.333a4.167 4.167 0 0 0 0-8.333m20.834 0a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func Delta(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M39.583 43.75L23.958 8.604zm2.084 0L25 6.25L8.333 43.75z"/>`), g.Group(children),
	)
}

func DeltaSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.958 16.667l9.375 18.75m2.084 0L25 14.583L14.583 35.417z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func DeskLampRound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 43.75V21.083A8.334 8.334 0 0 1 18.75 13L25 11.417M6.25 43.75h20.833m7.959-30.562l-2.084-4.521a4.17 4.17 0 1 0-7.562 3.52l2.083 4.417"/><path stroke="#344054" d="M30.917 14.187a8.333 8.333 0 0 1 11.075 4.03l1.76 3.777l-15.105 7.044l-1.76-3.776a8.33 8.33 0 0 1 4.03-11.075"/></g>`), g.Group(children),
	)
}

func DeskSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 12.5v25m33.334-25v25z"/><path stroke="#306CFE" d="M6.25 12.5h37.5M8.333 29.167h33.334V12.5H8.333z"/></g>`), g.Group(children),
	)
}

func DeskTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M41.667 20.833V37.5m-33.334-25v25z"/><path stroke="#306CFE" d="M6.25 12.5h37.5M25 20.833h16.667V12.5H25zm0 8.334h16.667v-8.334H25z"/></g>`), g.Group(children),
	)
}

func DiagramBarThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 18.75v12.5m-8.334-8.333v8.333zm-8.333-8.334V31.25z"/><path stroke="#306CFE" d="M10.417 6.25v37.5M6.25 39.583h37.5z"/></g>`), g.Group(children),
	)
}

func DiagramBarTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 35.417H14.583m4.167-25h-4.167zm4.167 8.333h-8.334zm4.166 8.333h-12.5z"/><path stroke="#306CFE" d="M6.25 6.25v35.417a2.083 2.083 0 0 0 2.083 2.083H43.75"/></g>`), g.Group(children),
	)
}

func Diameter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 27.083L35.417 25l-2.084-2.083m-16.666 0L14.583 25l2.084 2.083M14.583 25h20.834"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func DiameterCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25 24.896v.208"/><path stroke="#344054" stroke-width="2" d="M43.75 25H6.25"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M14.583 43.75H31.25zm12.5-31.25V8.333A2.083 2.083 0 0 0 25 6.25h-4.167a2.083 2.083 0 0 0-2.083 2.083V12.5zM18.75 25v18.75h8.333V25z"/><path stroke="#344054" d="M12.5 12.5a2.083 2.083 0 0 0-2.083 2.083v8.334A2.083 2.083 0 0 0 12.5 25h22.23a2.08 2.08 0 0 0 1.728-.937l2.771-4.167a2.08 2.08 0 0 0 0-2.292l-2.77-4.166a2.08 2.08 0 0 0-1.73-.938z"/></g>`), g.Group(children),
	)
}

func DirectionSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m43.23 17.354l-1.876-2.083a2.08 2.08 0 0 0-1.541-.688H29.167v8.334h10.645a2.08 2.08 0 0 0 1.542-.688l1.875-2.083a2.083 2.083 0 0 0 0-2.792M6.77 24.313l1.876 2.083a2.08 2.08 0 0 0 1.541.687h10.646V18.75H10.188a2.08 2.08 0 0 0-1.542.688L6.77 21.52a2.083 2.083 0 0 0 0 2.791"/><path stroke="#306CFE" d="M16.667 43.75h16.666m-4.166 0h-8.334V8.333a2.083 2.083 0 0 1 2.084-2.083h4.166a2.083 2.083 0 0 1 2.084 2.083z"/></g>`), g.Group(children),
	)
}

func Discount(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M30.104 30.208h.209M19.688 19.792h.208z"/><path stroke="#344054" stroke-width="2" d="m18.75 31.25l12.5-12.5"/><path stroke="#306CFE" stroke-width="2" d="M42.125 19.438c-.625-1.896-.208-4.334-1.354-5.896s-3.604-1.959-5.188-3.125C34 9.25 32.917 7.063 31.021 6.458C29.125 5.854 27.083 7 25 7s-4.167-1.146-6.02-.542c-1.855.605-2.98 2.834-4.563 3.959c-1.584 1.125-4 1.52-5.188 3.125s-.729 4-1.354 5.896S5.5 22.917 5.5 25s1.792 3.73 2.375 5.563s.208 4.333 1.354 5.895c1.146 1.563 3.604 1.959 5.188 3.105s2.666 3.375 4.562 3.979S22.917 43 25 43s4.167 1.146 6.02.542c1.855-.604 2.98-2.834 4.563-3.98c1.584-1.145 4.042-1.5 5.188-3.104s.729-4 1.354-5.895C42.75 28.667 44.5 27.083 44.5 25s-1.792-3.73-2.375-5.562"/></g>`), g.Group(children),
	)
}

func Diskette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 31.25H16.667v12.5h16.666zm-2.083-25h-12.5v8.333h12.5z"/><path stroke="#306CFE" d="M41.667 14.583v27.084a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083V8.333a2.083 2.083 0 0 1 2.084-2.083h22.916z"/></g>`), g.Group(children),
	)
}

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 8.333a4.167 4.167 0 1 0 0 8.333a4.167 4.167 0 0 0 0-8.333M29.167 37.5a4.167 4.167 0 1 0-8.334 0a4.167 4.167 0 0 0 8.334 0"/><path stroke="#306CFE" d="M6.25 25h37.5"/></g>`), g.Group(children),
	)
}

func DivideCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 15.625h.208m-.208 18.75h.208"/><path stroke="#344054" stroke-width="2" d="M14.583 25h20.834"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func DivideSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 15.625h.208m-.208 18.75h.208"/><path stroke="#344054" stroke-width="2" d="M14.583 25h20.834"/><path stroke="#306CFE" stroke-width="2" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Dome(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 27.083h-8.334v12.5h8.334z"/><path stroke="#306CFE" d="M6.25 39.583h37.5m-2.083 0H8.333v-12.5A16.667 16.667 0 0 1 25 10.417v0a16.667 16.667 0 0 1 16.667 16.666z"/></g>`), g.Group(children),
	)
}

func DoneCollection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 20.063l4.688 4.687l7.812-7.812"/><path stroke="#306CFE" d="M33.333 6.25h-25c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.084 2.083 2.084h25c1.15 0 2.084-.933 2.084-2.084v-25c0-1.15-.933-2.083-2.084-2.083"/><path stroke="#306CFE" d="M14.583 43.75h27.084a2.083 2.083 0 0 0 2.083-2.083v-31.25"/></g>`), g.Group(children),
	)
}

func DonePlaylistTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m35.417 11.292l2.77 2.937l5.563-5.896m-14.583 31.25H43.75M29.167 31.25H43.75z"/><path stroke="#306CFE" d="M20.833 36.458V6.25zm-7.291-7.291a7.292 7.292 0 1 0 0 14.583a7.292 7.292 0 0 0 0-14.583"/></g>`), g.Group(children),
	)
}

func DoorHandle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="M29.167 22.917v10.416a10.416 10.416 0 1 1-20.834 0V16.667A10.417 10.417 0 0 1 18.75 6.25v0a10.42 10.42 0 0 1 10.208 8.333"/><path stroke="#344054" stroke-width="2" d="M37.5 22.917H22.917a4.167 4.167 0 1 1 0-8.334H37.5a4.167 4.167 0 1 1 0 8.334"/><path stroke="#344054" stroke-width="3" d="M18.646 33.333h.208"/></g>`), g.Group(children),
	)
}

func DoubleDownSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 10.417L25 20.833L6.25 10.417"/><path stroke="#306CFE" d="M43.75 29.167L25 39.583L6.25 29.167"/></g>`), g.Group(children),
	)
}

func DoubleDownSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 27.083L25 33.333l-6.25-6.25m12.5-10.416L25 22.917l-6.25-6.25"/><path stroke="#306CFE" d="M43.75 41.667V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083"/></g>`), g.Group(children),
	)
}

func DoubleLeftSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 6.25L29.167 25l10.416 18.75"/><path stroke="#306CFE" d="M20.833 6.25L10.417 25l10.416 18.75"/></g>`), g.Group(children),
	)
}

func DoubleLeftSignCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M21.875 18.75L15.625 25l6.25 6.25m10.417-12.5L26.042 25l6.25 6.25"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func DoubleLeftSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 31.25L16.667 25l6.25-6.25m10.416 12.5L27.083 25l6.25-6.25"/><path stroke="#306CFE" d="M8.333 43.75h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083"/></g>`), g.Group(children),
	)
}

func DoubleRightSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 43.75L20.833 25L10.417 6.25"/><path stroke="#306CFE" d="M29.167 43.75L39.583 25L29.167 6.25"/></g>`), g.Group(children),
	)
}

func DoubleRightSignCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m28.125 31.25l6.25-6.25l-6.25-6.25m-10.417 12.5l6.25-6.25l-6.25-6.25"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func DoubleRightSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m27.083 18.75l6.25 6.25l-6.25 6.25m-10.416-12.5l6.25 6.25l-6.25 6.25"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func DoubleUpScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.5 25h-25"/><path stroke="#306CFE" d="M20.833 10.417L25 6.25l4.167 4.167M20.833 37.5L25 33.333l4.167 4.167M25 33.333V43.75m0-27.083V6.25z"/></g>`), g.Group(children),
	)
}

func DoubleUpSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 39.583L25 29.167L6.25 39.583"/><path stroke="#306CFE" d="M43.75 20.833L25 10.417L6.25 20.833"/></g>`), g.Group(children),
	)
}

func DoubleUpSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.75 22.917l6.25-6.25l6.25 6.25m-12.5 10.416l6.25-6.25l6.25 6.25"/><path stroke="#306CFE" d="M6.25 8.333v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083"/></g>`), g.Group(children),
	)
}

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 6.25v37.5"/><path stroke="#344054" d="M18.75 37.5L25 43.75l6.25-6.25"/></g>`), g.Group(children),
	)
}

func DownDirectionTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.146 13.104L25 25.958l12.854-12.854"/><path stroke="#306CFE" d="m37.854 13.104l5.896 5.875l-17.27 17.292a2.08 2.08 0 0 1-2.96 0L6.25 19l5.896-5.897"/></g>`), g.Group(children),
	)
}

func DownJunctionSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 6.25H16.667m25 8.333h2.083zm0-8.333h2.083zM8.333 14.583H6.25zm0-8.333H6.25z"/><path stroke="#306CFE" d="M16.667 14.583h4.166V31.25h-6.25l8.813 10.583a2.083 2.083 0 0 0 3.208 0l8.813-10.583h-6.25V14.583h4.166"/></g>`), g.Group(children),
	)
}

func DownOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 16.667v16.666m5.208-5.208L25 33.333l-5.208-5.208"/><path stroke="#306CFE" d="M6.25 32.77V17.23L17.23 6.25h15.54l10.98 10.98v15.54L32.77 43.75H17.23z"/></g>`), g.Group(children),
	)
}

func DownTrend(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m6.25 12.5l16.667 16.667l6.25-6.25L43.75 37.5"/><path stroke="#344054" d="M35.417 37.5h8.333v-8.333"/></g>`), g.Group(children),
	)
}

func DownUpScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.5 25h-25"/><path stroke="#306CFE" d="M20.833 37.5L25 33.333l4.167 4.167m0-25L25 16.667L20.833 12.5M25 33.333V43.75m0-37.5v10.417z"/></g>`), g.Group(children),
	)
}

func DownloadAltFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25v16.667m-6.25-6.25l6.25 6.25l6.25-6.25"/><path stroke="#306CFE" d="M43.75 27.083L37.98 9.75A2.08 2.08 0 0 0 36 8.333h-2.667m-16.666 0H14a2.08 2.08 0 0 0-2.083 1.417L6.25 27.083m37.5 0v14.584a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V27.083h10.417a8.333 8.333 0 1 0 16.666 0z"/></g>`), g.Group(children),
	)
}

func DownloadAltFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 25v18.75m-6.48-6.25l6.626 6.625l6.333-6.333"/><path stroke="#306CFE" d="M12.5 31.48a13.73 13.73 0 1 1 20.833-15.084a9.5 9.5 0 0 1 1.73-.167A8.729 8.729 0 0 1 37.5 33.333"/></g>`), g.Group(children),
	)
}

func DownloadFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 29.167L25 35.417l-6.25-6.25m6.25 6.25V6.25"/><path stroke="#306CFE" d="M8.333 35.417v6.25a2.083 2.083 0 0 0 2.084 2.083h29.166a2.083 2.083 0 0 0 2.084-2.083v-6.25"/></g>`), g.Group(children),
	)
}

func DrawersThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 20.833V25m-16.666 0v-4.167z"/><path stroke="#306CFE" d="M43.75 6.25H6.25m4.167 37.5h29.166a2.083 2.083 0 0 0 2.084-2.083V6.25H8.333v35.417a2.083 2.083 0 0 0 2.084 2.083m0 0H25V6.25H8.333v35.417a2.083 2.083 0 0 0 2.084 2.083"/></g>`), g.Group(children),
	)
}

func DrawersTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 37.5v4.167M10.417 37.5v4.167z"/><path stroke="#306CFE" d="M6.25 8.333h37.5M10.417 37.5h29.166a2.083 2.083 0 0 0 2.084-2.083v-18.75H8.333v18.75a2.083 2.083 0 0 0 2.084 2.083m0 0H25V16.667H8.333v18.75a2.083 2.083 0 0 0 2.084 2.083M8.333 16.667h33.334V8.333H8.333z"/></g>`), g.Group(children),
	)
}

func DrawingTabletPencil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 16.667h6.25a2.083 2.083 0 0 1 2.083 2.083v22.917a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V18.75a2.083 2.083 0 0 1 2.083-2.083h14.584"/><path stroke="#306CFE" d="M8.333 16.667A2.083 2.083 0 0 0 6.25 18.75v22.917a2.083 2.083 0 0 0 2.083 2.083h6.25V16.667zM38.98 9.77l-2.916-2.917a2.083 2.083 0 0 0-2.917 0l-10.23 10.23v5.833h5.834l10.23-10.23a2.083 2.083 0 0 0 0-2.916"/></g>`), g.Group(children),
	)
}

func DumbbellThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 20.02l-1.979-8.708a4.167 4.167 0 0 1 4.063-5.062h16.666a4.167 4.167 0 0 1 4.167 5.062l-2.083 8.75M18.75 35.417h12.5"/><path stroke="#306CFE" d="M35.792 43.75H14.208a4.17 4.17 0 0 1-3.562-1.958a16.666 16.666 0 1 1 28.687 0a4.17 4.17 0 0 1-3.541 1.958"/></g>`), g.Group(children),
	)
}

func Earphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.5 18.75v25m-25 0v-25z"/><path stroke="#306CFE" d="M14.583 18.75h-4.166a2.083 2.083 0 0 1-2.084-2.083V8.333a2.083 2.083 0 0 1 2.084-2.083h4.166a6.25 6.25 0 0 1 6.25 6.25v0a6.25 6.25 0 0 1-6.25 6.25m20.834 0h4.166a2.083 2.083 0 0 0 2.084-2.083V8.333a2.083 2.083 0 0 0-2.084-2.083h-4.166a6.25 6.25 0 0 0-6.25 6.25v0a6.25 6.25 0 0 0 6.25 6.25"/></g>`), g.Group(children),
	)
}

func EarphoneBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M13.688 26.604a14.4 14.4 0 0 1-1.188-5.77a14.583 14.583 0 1 1 29.167 0M18.75 33.333l-2.083 2.084"/><path stroke="#306CFE" d="M35.82 28.042L22.56 41.3a8.333 8.333 0 0 1-11.785-11.785l13.259-13.259a2.083 2.083 0 0 1 2.946 0l8.839 8.84a2.083 2.083 0 0 1 0 2.946"/></g>`), g.Group(children),
	)
}

func EditAltFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.667 17.167L15.083 43.75H6.25v-8.833L32.833 8.333a6.25 6.25 0 0 1 8.834 0v0a6.25 6.25 0 0 1 0 8.834"/>`), g.Group(children),
	)
}

func EditAltSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M42.52 13.354L36.647 7.48a2.083 2.083 0 0 0-2.959 0l-6 6l8.834 8.834l6-6a2.084 2.084 0 0 0 0-2.959"/><path stroke="#306CFE" d="m21.813 19.354l8.833 8.834L15.083 43.75H6.25v-8.833z"/></g>`), g.Group(children),
	)
}

func EditCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m28.708 27.188l-5.791-.105v-5.791L36.312 7.896a2.083 2.083 0 0 1 2.938 0l2.958 2.958a2.083 2.083 0 0 1 0 2.938z"/><path stroke="#306CFE" d="M43.75 25A18.75 18.75 0 1 1 25 6.25"/></g>`), g.Group(children),
	)
}

func EditUserSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m43.146 33.52l-10.23 10.23h-5.833v-5.833l10.23-10.23a2.084 2.084 0 0 1 2.916 0l2.917 2.917a2.084 2.084 0 0 1 0 2.917"/><path stroke="#306CFE" d="M35.417 19.313q.015-.282 0-.563A12.5 12.5 0 1 0 22 31.25"/><path stroke="#306CFE" d="M18.75 43.75c-9.27-.77-12.5-4.167-12.5-4.167V37.5a10.42 10.42 0 0 1 7.27-9.896"/></g>`), g.Group(children),
	)
}

func EmailFile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 6.25v4.167H37.5zm-9.479 24.23a2.08 2.08 0 0 0 2.292 0L37.5 22.916v-12.5L33.333 6.25h-18.75A2.083 2.083 0 0 0 12.5 8.333v14.584z"/><path stroke="#306CFE" d="m26.146 30.48l14.375-9.647a2.083 2.083 0 0 1 3.229 1.73v19.104a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V22.646a2.084 2.084 0 0 1 3.23-1.813l14.374 9.563a2.08 2.08 0 0 0 2.292.083"/></g>`), g.Group(children),
	)
}

func EmailOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M23.854 30.48a2.08 2.08 0 0 0 2.292 0L37.5 22.916V6.25h-25v16.667z"/><path stroke="#306CFE" d="m26.146 30.48l14.375-9.647a2.083 2.083 0 0 1 3.229 1.73v19.104a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V22.646a2.084 2.084 0 0 1 3.23-1.813l14.374 9.563a2.08 2.08 0 0 0 2.292.083"/></g>`), g.Group(children),
	)
}

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 18.75l9.125 7.27a2.08 2.08 0 0 0 2.584 0l9.125-7.27"/><path stroke="#306CFE" d="M41.667 10.417H8.333c-1.15 0-2.083.932-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Exchange(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 35.417l-6.25-6.25h33.334"/><path stroke="#306CFE" d="m35.417 14.583l6.25 6.25H8.333"/></g>`), g.Group(children),
	)
}

func Exclamation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M27.083 6.25h-4.166c-1.15 0-2.084.933-2.084 2.083V25c0 1.15.933 2.083 2.084 2.083h4.166c1.15 0 2.084-.932 2.084-2.083V8.333c0-1.15-.933-2.083-2.084-2.083"/><path stroke="#344054" d="M25 43.75a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/></g>`), g.Group(children),
	)
}

func ExportTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 6.25h12.5v12.5m-20.833 8.333L43.75 6.25"/><path stroke="#306CFE" d="M43.75 27.083v14.584a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h14.584"/></g>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 25S33.063 33.333 25 33.333C16.938 33.333 10.417 25 10.417 25s6.52-8.333 14.583-8.333S39.583 25 39.583 25M25 20.833a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/><path stroke="#306CFE" d="M6.25 14.583v-6.25A2.083 2.083 0 0 1 8.333 6.25h6.25m29.167 8.333v-6.25a2.083 2.083 0 0 0-2.083-2.083h-6.25M6.25 35.417v6.25a2.083 2.083 0 0 0 2.083 2.083h6.25m29.167-8.333v6.25a2.083 2.083 0 0 1-2.083 2.083h-6.25"/></g>`), g.Group(children),
	)
}

func EyeAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M35.417 25S30.75 31.25 25 31.25S14.583 25 14.583 25S19.25 18.75 25 18.75S35.417 25 35.417 25"/><path stroke="#344054" stroke-width="3" d="M25.104 25h-.208"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m25 25l18.75 12.5v-25z"/><path stroke="#306CFE" d="M25 12.5v25L6.25 25z"/></g>`), g.Group(children),
	)
}

func FastBackwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 25l8.334 6.25v-12.5zm8.334 0l8.333 6.25v-12.5z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 12.5v25L25 25z"/><path stroke="#306CFE" d="M43.75 25L25 37.5v-25z"/></g>`), g.Group(children),
	)
}

func FastForwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 18.75v12.5L35.417 25zm-8.333 0v12.5L27.083 25z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Favourite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M40.77 11.333a10.23 10.23 0 0 1 0 14.438L25 41.667L9.23 25.77a10.229 10.229 0 0 1 7.166-17.438a10.2 10.2 0 0 1 7.166 3A9.3 9.3 0 0 1 25 13.167a9.3 9.3 0 0 1 1.438-1.834a10.06 10.06 0 0 1 14.333 0"/>`), g.Group(children),
	)
}

func FavouriteAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.75 18.75c0 17.042-18.75 25-18.75 25s-18.75-7.875-18.75-25a9.98 9.98 0 0 1 9.375-10.417A9.98 9.98 0 0 1 25 18.75a9.98 9.98 0 0 1 9.375-10.417A9.98 9.98 0 0 1 43.75 18.75"/>`), g.Group(children),
	)
}

func FileFavoriteEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 18.75h12.5m8.334 10.417l-2.584 4.791l-5.75.771l4.167 3.75l-.98 5.271l5.147-2.5l5.145 2.5l-.979-5.27l4.167-3.75l-5.75-.772zm-20.834-2.084h12.5z"/><path stroke="#306CFE" d="M20.833 43.75h-12.5a2.083 2.083 0 0 1-2.083-2.083V14.583l8.333-8.333h18.75a2.083 2.083 0 0 1 2.084 2.083v12.5"/></g>`), g.Group(children),
	)
}

func FileFolderApprovedTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m31.25 39.583l4.167 4.167l8.333-8.333M35.417 6.25H12.5a2.083 2.083 0 0 0-2.083 2.083V18.75h8.333l4.167 4.167h16.666v-12.5zm0 0v4.167h4.166z"/><path stroke="#306CFE" d="M22.917 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083V20.833a2.083 2.083 0 0 1 2.083-2.083H18.75l4.167 4.167h18.75A2.083 2.083 0 0 1 43.75 25v2.083"/></g>`), g.Group(children),
	)
}

func FileMusicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 33.333a6.04 6.04 0 0 0-6.25-6.25v12.5"/><path stroke="#344054" d="M29.167 43.75a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/><path stroke="#306CFE" d="M39.583 18.75V8.333A2.083 2.083 0 0 0 37.5 6.25H14.583l-4.166 4.167v31.25A2.083 2.083 0 0 0 12.5 43.75h4.167"/><path stroke="#306CFE" d="M10.417 10.417h4.166V6.25z"/></g>`), g.Group(children),
	)
}

func FileSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 18.75h12.5m0 16.667h-12.5zm-12.5-8.334h12.5z"/><path stroke="#306CFE" d="M39.583 10.417v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 12.5 6.25h22.917zM35.417 6.25v4.167h4.166z"/></g>`), g.Group(children),
	)
}

func FileThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 12.5 6.25h14.583l12.5 12.5v22.917A2.083 2.083 0 0 1 37.5 43.75"/><path stroke="#344054" d="m39.583 18.75l-12.5-12.5L25 16.667z"/></g>`), g.Group(children),
	)
}

func FileVideoFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m27.083 25l-4.166-3.125v6.25z"/><path stroke="#306CFE" d="M12.5 43.75a2.083 2.083 0 0 1-2.083-2.083V14.583L18.75 6.25H37.5a2.083 2.083 0 0 1 2.083 2.083v33.334A2.083 2.083 0 0 1 37.5 43.75z"/></g>`), g.Group(children),
	)
}

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 22.917h4.167M20.833 6.25v20.833zm-2.083 8.333h4.167z"/><path stroke="#306CFE" d="M39.583 10.417v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 12.5 6.25h22.917zM35.417 6.25v4.167h4.166z"/></g>`), g.Group(children),
	)
}

func Files(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M16.77 35.417h-.207m16.875 0h-.209"/><path stroke="#344054" stroke-width="2" d="M33.333 12.5v8.333M16.667 12.5v8.333z"/><path stroke="#306CFE" stroke-width="2" d="M22.917 45.833h-12.5a2.083 2.083 0 0 1-2.084-2.083V6.25a2.083 2.083 0 0 1 2.084-2.083h12.5A2.083 2.083 0 0 1 25 6.25v37.5a2.083 2.083 0 0 1-2.083 2.083m18.75-2.083V6.25a2.083 2.083 0 0 0-2.084-2.083h-12.5A2.083 2.083 0 0 0 25 6.25v37.5a2.083 2.083 0 0 0 2.083 2.083h12.5a2.083 2.083 0 0 0 2.084-2.083"/></g>`), g.Group(children),
	)
}

func FilmMovie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 35.417v8.333h12.5v-8.333zm-12.5 0v8.333H25v-8.333zM6.25 43.75h37.5zM37.5 14.583V6.25H25v8.333zm-12.5 0V6.25H12.5v8.333zM6.25 6.25h37.5zM27.083 25l-4.166-3.125v6.25z"/><path stroke="#306CFE" d="M43.75 33.333V16.667a2.084 2.084 0 0 0-2.083-2.084H8.333c-1.15 0-2.083.933-2.083 2.084v16.666c0 1.15.933 2.084 2.083 2.084h33.334c1.15 0 2.083-.933 2.083-2.084"/></g>`), g.Group(children),
	)
}

func FilmRoll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M22.813 17.708h.208M15.52 25h.21m7.29 7.292h-.207M30.104 25h.209"/><path stroke="#306CFE" stroke-width="2" d="M37.333 33.333a16.67 16.67 0 0 1-14.416 8.334H43.75a2.083 2.083 0 0 0 2.083-2.084v-4.166a2.083 2.083 0 0 0-2.083-2.084zM39.583 25a16.666 16.666 0 1 1-33.332 0a16.666 16.666 0 0 1 33.332 0"/></g>`), g.Group(children),
	)
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.417 8.333v5.5c-.001.487.168.959.479 1.334L20.354 26.5c.31.375.48.847.48 1.333V43.75l8.333-4.167v-11.75a2.08 2.08 0 0 1 .479-1.333l9.458-11.333c.31-.375.48-.847.48-1.334v-5.5A2.083 2.083 0 0 0 37.5 6.25h-25a2.083 2.083 0 0 0-2.083 2.083"/>`), g.Group(children),
	)
}

func FilterAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 20.833V6.25m-29.166 0v29.167zM25 14.583V43.75zm14.583 14.584V43.75z"/><path stroke="#344054" d="M10.417 35.417a4.166 4.166 0 1 0 0 8.332a4.166 4.166 0 0 0 0-8.332M25 6.25a4.167 4.167 0 1 0 0 8.333a4.167 4.167 0 0 0 0-8.333m14.583 14.583a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func FilterAltFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 25h29.166"/><path stroke="#306CFE" d="M14.583 35.417h20.834M6.25 14.583h37.5z"/></g>`), g.Group(children),
	)
}

func FilterAltFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.27 9.708l11.563 13.209V43.75l8.334-4.167V22.917L40.729 9.708a2.083 2.083 0 0 0-1.562-3.458H10.833a2.083 2.083 0 0 0-1.562 3.458"/>`), g.Group(children),
	)
}

func FireLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25 43.75c8.125 0 14.583-4.167 14.583-14.583S29.167 18.75 29.167 6.25c-6.25 4.167-9.105 8.542-10.417 16.667a10.42 10.42 0 0 1-4.167-6.25C12.5 18.75 10.417 25 10.417 29.167c0 6.541 2.666 14.583 14.583 14.583"/>`), g.Group(children),
	)
}

func Fireplace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 14.583h16.666zm14.812 22.688a10.8 10.8 0 0 0-1.333-5.563A4.5 4.5 0 0 1 27.77 34.5a9.6 9.6 0 0 0-4.604-7.417c0 5.563-4.646 5.563-4.646 10.188A6.354 6.354 0 0 0 25 43.75c5.563 0 6.48-3.563 6.48-6.48"/><path stroke="#306CFE" d="M6.25 6.25h37.5m-2.083 37.5H8.333V6.25h33.334zm-35.417 0h37.5z"/></g>`), g.Group(children),
	)
}

func FishingHook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m31.25 33.333l4.167-2.083v4.167a8.333 8.333 0 0 1-16.667 0V20.833m0-8.333V6.25"/><path stroke="#344054" d="M18.75 20.833a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/></g>`), g.Group(children),
	)
}

func FlaskThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 43.75h-25a2.084 2.084 0 0 1-1.833-3.062l10.166-18.896V6.25h8.334v15.542l10.166 18.896A2.083 2.083 0 0 1 37.5 43.75"/><path stroke="#344054" d="M18.75 6.25h12.5"/></g>`), g.Group(children),
	)
}

func FlaskTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 31.25a12.5 12.5 0 1 1-16.667-11.77V6.25h8.334v13.23A12.5 12.5 0 0 1 37.5 31.25"/><path stroke="#344054" d="M18.75 6.25h12.5"/></g>`), g.Group(children),
	)
}

func FloorLamp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 43.75h8.334m-4.167 0V20.833z"/><path stroke="#306CFE" d="M35.417 20.833a28.3 28.3 0 0 1-4.063-12.708a2.083 2.083 0 0 0-2.083-1.875h-8.563a2.083 2.083 0 0 0-2.083 1.875a28.3 28.3 0 0 1-4.042 12.708z"/></g>`), g.Group(children),
	)
}

func FloppyDiskAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M25 27.083v6.25"/><path stroke="#344054" stroke-width="3" d="M25.104 42.708h-.208"/><path stroke="#306CFE" stroke-width="2" d="M35.417 43.75h4.166a2.083 2.083 0 0 0 2.084-2.083V14.583L33.333 6.25H10.417a2.083 2.083 0 0 0-2.084 2.083v33.334a2.083 2.083 0 0 0 2.084 2.083h4.166"/><path stroke="#306CFE" stroke-width="2" d="M31.25 6.25h-12.5v8.333h12.5z"/></g>`), g.Group(children),
	)
}

func FootballBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.146 6.896L25 10.416l4.854-3.52M6.292 24.02l4.833-3.52l-1.833-5.708M37.5 37.5l-4-.083l-1.52 4.958M12.5 37.5l4-.083l1.52 4.958m22.71-27.583L38.874 20.5l4.833 3.52M25 16.667v-6.25zm-7.48 5.52L11.126 20.5zm2.98 9.063l-4.062 6.25zm9 0l4.063 6.25zm2.792-9.02l6.583-1.73zM25 16.666l-7.292 5.562L20.5 31.25h9l2.792-9.02z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func ForestTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M27.083 17.708a7.27 7.27 0 0 0-4.354-6.666a6.25 6.25 0 0 0-12.125 0a7.27 7.27 0 0 0 0 13.333a6.25 6.25 0 0 0 12.125 0a7.27 7.27 0 0 0 4.354-6.667"/><path stroke="#344054" d="M16.667 18.75v25m22.916-12.5v12.5"/><path stroke="#306CFE" d="M6.25 43.75h37.5m0-16.667v-6.25a4.167 4.167 0 0 0-8.333 0v6.25a4.167 4.167 0 1 0 8.333 0"/></g>`), g.Group(children),
	)
}

func Forklift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 37.5h-8.333v-25m-25 25H8.333a2.083 2.083 0 0 1-2.083-2.083v-12.5a2.083 2.083 0 0 1 2.083-2.084h7.042a2.08 2.08 0 0 1 1.875 1.146L18.75 25h16.667v12.5m-8.334 0h-8.166"/><path stroke="#306CFE" d="M10.417 20.583V10.417A2.083 2.083 0 0 1 12.5 8.333h8.792a2.08 2.08 0 0 1 2.083 1.584l3.708 14.812"/><path stroke="#344054" d="M31.25 33.333a4.166 4.166 0 1 1 0 8.333a4.166 4.166 0 0 1 0-8.333M10.417 37.5a4.167 4.167 0 1 0 8.334 0a4.167 4.167 0 0 0-8.334 0"/></g>`), g.Group(children),
	)
}

func ForwardTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 43.75V6.25"/><path stroke="#306CFE" d="M10.417 10.417v29.166L31.25 25z"/></g>`), g.Group(children),
	)
}

func FourKTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.708 23.313l10.042 16.27M29.167 10.417v29.166zm0 18.75l14.583-18.75z"/><path stroke="#306CFE" d="M20.833 10.417L6.25 29.167h14.583m0 10.416V10.417z"/></g>`), g.Group(children),
	)
}

func FryingPan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.458 28.688L41.5 34.25a5.207 5.207 0 0 1-1.197 9.17a5.21 5.21 0 0 1-6.053-1.92l-5.583-8.062"/><path stroke="#306CFE" d="M20.833 35.417c8.055 0 14.584-6.53 14.584-14.584S28.887 6.25 20.833 6.25S6.25 12.78 6.25 20.833c0 8.055 6.53 14.584 14.583 14.584"/></g>`), g.Group(children),
	)
}

func FullCrossCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.75 31.25l12.5-12.5m-12.5 0l12.5 12.5"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Gallery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M18.146 20.833L6.25 32.771v6.812a2.084 2.084 0 0 0 2.083 2.084h8.938l10.833-10.834zM33.938 25L17.27 41.667h24.396a2.083 2.083 0 0 0 2.083-2.084v-4.77z"/><path stroke="#344054" stroke-width="3" d="M29 17.27h-.208"/><path stroke="#306CFE" stroke-width="2" d="M41.667 8.333H8.333c-1.15 0-2.083.933-2.083 2.084v29.166c0 1.15.933 2.084 2.083 2.084h33.334c1.15 0 2.083-.933 2.083-2.084V10.417a2.084 2.084 0 0 0-2.083-2.084"/></g>`), g.Group(children),
	)
}

func GalleryCollections(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 41.667h31.25a2.083 2.083 0 0 0 2.083-2.084V16.667"/><path stroke="#306CFE" d="m23.417 25.354l-7.98 7.98H8.334A2.083 2.083 0 0 1 6.25 31.25V28l9.917-9.917zm5.916-5.917L15.438 33.333h17.895a2.083 2.083 0 0 0 2.084-2.083v-5.708zm-21-11.104h25a2.083 2.083 0 0 1 2.084 2.084V31.25a2.083 2.083 0 0 1-2.084 2.083h-25A2.083 2.083 0 0 1 6.25 31.25V10.417a2.083 2.083 0 0 1 2.083-2.084"/></g>`), g.Group(children),
	)
}

func GameConsoleCable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25v8.333m6.25 0h-12.5l1.042 4.167h10.416z"/><path stroke="#306CFE" d="M34.23 14.583a8.333 8.333 0 0 1 8.333 7.75L43.75 39a4.418 4.418 0 0 1-4.667 4.75a4.67 4.67 0 0 1-4.166-3.792l-1.584-7.062a2.08 2.08 0 0 0-2.083-1.646H18.625a2.084 2.084 0 0 0-2.083 1.646l-1.521 7.062a4.67 4.67 0 0 1-4.167 3.792A4.416 4.416 0 0 1 6.25 39l1.188-16.667a8.334 8.334 0 0 1 8.333-7.75z"/></g>`), g.Group(children),
	)
}

func GasStove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="m32.792 17.208l2.625-2.625m-3.042 17.792a10.416 10.416 0 1 1-14.713-14.75a10.416 10.416 0 0 1 14.713 14.75m.438.437l2.604 2.605zM17.27 32.73l-2.688 2.688zm-.063-15.52l-2.625-2.626z"/><path stroke="#344054" stroke-width="3" d="M24.896 25h.208"/><path stroke="#306CFE" stroke-width="2" d="M43.75 41.667V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083"/></g>`), g.Group(children),
	)
}

func GearshiftCar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 31.25v-12.5M33.333 25H16.667zM25 31.25v-12.5zm8.333 0v-12.5z"/><path stroke="#306CFE" d="M10.417 43.75h29.166c1.15 0 2.084-.933 2.084-2.083V8.333c0-1.15-.933-2.083-2.084-2.083H10.417a2.084 2.084 0 0 0-2.084 2.083v33.334c0 1.15.933 2.083 2.084 2.083"/></g>`), g.Group(children),
	)
}

func Goal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m31.25 18.75l-5.208 5.208zm0-6.25v6.25h6.25l6.25-6.25H37.5V6.25z"/><path stroke="#306CFE" d="M25.688 6.25H25A18.75 18.75 0 1 0 43.75 25v-.687"/><path stroke="#306CFE" d="M35.208 27.083a10.417 10.417 0 1 1-12.291-12.291"/></g>`), g.Group(children),
	)
}

func GoldMedal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.708 17.5l-2.875.396l2.084 1.854l-.5 2.646L25 21.146l2.583 1.25l-.5-2.646l2.084-1.854l-2.875-.396L25 15.104z"/><path stroke="#306CFE" d="M25 31.25c6.904 0 12.5-5.596 12.5-12.5S31.904 6.25 25 6.25s-12.5 5.596-12.5 12.5s5.596 12.5 12.5 12.5"/><path stroke="#306CFE" d="M33.917 27.5a12.5 12.5 0 0 1-6.834 3.583l4.167 10.584l3.125-2.605l4.354.438zM11.27 39.583l4.355-.437l3.125 2.52l4.167-10.583a12.5 12.5 0 0 1-6.834-3.583z"/></g>`), g.Group(children),
	)
}

func GoldMedalSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 27.083a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/><path stroke="#306CFE" d="M12.063 28.27L6.25 36.584l6.313 1.875l3.916 5.292l5.875-8.375m15.584-7.106l5.812 8.313l-6.312 1.875l-3.917 5.292l-5.875-8.375"/><path stroke="#306CFE" d="M40.167 20.833c0 1.584-1.396 2.917-1.855 4.334c-.458 1.416-.145 3.354-1.041 4.583s-2.813 1.5-3.938 2.417s-2.083 2.625-3.541 3.104c-1.459.479-3.105-.438-4.688-.438s-3.27.896-4.687.438c-1.417-.459-2.313-2.23-3.542-3.104c-1.23-.875-3.146-1.188-4.042-2.417c-.896-1.23-.562-3.125-1.041-4.583c-.48-1.459-1.959-2.75-1.959-4.334s1.396-2.916 1.854-4.333s.146-3.354 1.042-4.583s2.813-1.5 4.042-2.417s1.979-2.625 3.541-3.104c1.563-.48 3.105.437 4.688.437s3.27-.896 4.688-.437C31.104 6.854 32 8.625 33.333 9.5s3.146 1.187 4.042 2.417s.563 3.125 1.042 4.583s1.75 2.75 1.75 4.333"/></g>`), g.Group(children),
	)
}

func GpsFixed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25v4.167zM43.75 25h-4.167zM25 43.75v-4.167zM6.25 25h4.167zm25 0a6.25 6.25 0 1 0-12.5 0a6.25 6.25 0 0 0 12.5 0"/><path stroke="#306CFE" d="M25 39.583c8.054 0 14.583-6.529 14.583-14.583S33.054 10.417 25 10.417S10.417 16.946 10.417 25S16.946 39.583 25 39.583"/></g>`), g.Group(children),
	)
}

func Gravity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 8.333v12.5M25 6.25v12.5zm8.333 2.083v12.5z"/><path stroke="#306CFE" d="M25 43.75a8.333 8.333 0 1 0 0-16.667a8.333 8.333 0 0 0 0 16.667"/></g>`), g.Group(children),
	)
}

func Grill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 35.417h25"/><path stroke="#344054" d="m35.125 23.667l4.458 20.083m-29.166 0l4.458-20.02"/><path stroke="#306CFE" d="M41.667 6.25v4.167a16.667 16.667 0 0 1-33.334 0V6.25z"/><path stroke="#344054" d="M6.25 6.25h37.5"/></g>`), g.Group(children),
	)
}

func GrillBbq(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.333 25l3.542 10.417M10.417 43.75l4.458-20.083"/><path stroke="#306CFE" d="M41.667 6.25v4.167a16.667 16.667 0 1 1-33.334 0V6.25z"/><path stroke="#344054" d="M6.25 6.25h37.5"/><path stroke="#306CFE" d="M37.5 43.75a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/></g>`), g.Group(children),
	)
}

func HammerDrill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 10.417H8.333M25 31.25v12.5zm8.333-20.833h8.334z"/><path stroke="#306CFE" d="M31.25 22.917h-12.5a2.083 2.083 0 0 1-2.083-2.084v-12.5A2.083 2.083 0 0 1 18.75 6.25h12.5a2.083 2.083 0 0 1 2.083 2.083v12.5a2.083 2.083 0 0 1-2.083 2.084m-2.083 0h-8.334v6.25a2.083 2.083 0 0 0 2.084 2.083h4.166a2.083 2.083 0 0 0 2.084-2.083z"/></g>`), g.Group(children),
	)
}

func Handphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 35.417h.208"/><path stroke="#344054" stroke-width="2" d="M29.167 6.25h-8.334l1.042 4.167h6.25z"/><path stroke="#306CFE" stroke-width="2" d="M35.417 6.25H14.583c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h20.834c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func HandphoneLaptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 6.25H10.417a2.083 2.083 0 0 0-2.084 2.083V18.75h14.584A2.083 2.083 0 0 1 25 20.833V31.25h16.667V8.333a2.083 2.083 0 0 0-2.084-2.083M25 31.25v4.167h16.667a2.083 2.083 0 0 0 2.083-2.084V31.25z"/><path stroke="#344054" d="M6.25 41.667V20.833a2.083 2.083 0 0 1 2.083-2.083h14.584A2.083 2.083 0 0 1 25 20.833v20.834a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083m7.292-18.75h4.166l1.042-4.167H12.5z"/></g>`), g.Group(children),
	)
}

func HandphoneLock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 43.75h12.5a2.083 2.083 0 0 0 2.084-2.083v-8.334a2.083 2.083 0 0 0-2.084-2.083h-12.5A2.083 2.083 0 0 0 25 33.333v8.334a2.083 2.083 0 0 0 2.083 2.083m2.084-16.667v4.167H37.5v-4.167a4.167 4.167 0 0 0-4.167-4.166v0a4.167 4.167 0 0 0-4.166 4.166"/><path stroke="#306CFE" d="M16.667 43.75h-6.25a2.083 2.083 0 0 1-2.084-2.083V8.333a2.083 2.083 0 0 1 2.084-2.083H31.25a2.083 2.083 0 0 1 2.083 2.083v6.25"/><path stroke="#306CFE" d="M17.708 10.417h6.25L25 6.25h-8.333z"/></g>`), g.Group(children),
	)
}

func Hastag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 31.25h31.25m-29.166-12.5h31.25z"/><path stroke="#306CFE" d="m13.5 43.75l10.063-37.5m12.937 0l-10.062 37.5z"/></g>`), g.Group(children),
	)
}

func HastagCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.48 14.583l-5.022 18.75m-1.791-12.5h18.75zm-2.084 8.334h18.75zm11.938 6.25l5.02-18.75z"/><path stroke="#306CFE" d="M6.25 25A18.75 18.75 0 0 1 25 6.25v0A18.75 18.75 0 0 1 43.75 25v0A18.75 18.75 0 0 1 25 43.75v0A18.75 18.75 0 0 1 6.25 25"/></g>`), g.Group(children),
	)
}

func HastagSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.48 14.583l-5.584 20.834m-1.23-14.584h18.75zm-2.083 8.334h18.75zm17.521-14.584l-5.583 20.834z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func HeadsetAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 29.167L7.938 25A17.146 17.146 0 0 1 25 6.25v0A17.146 17.146 0 0 1 42.063 25l-.396 4.167m-6.25 6.25V37.5a3.94 3.94 0 0 1-3.646 4.167H25"/><path stroke="#306CFE" d="M8.333 29.167v-2.084a6.25 6.25 0 0 1 6.25-6.25h2.084a2.083 2.083 0 0 1 2.083 2.084v10.416a2.083 2.083 0 0 1-2.083 2.084h-2.084a6.25 6.25 0 0 1-6.25-6.25m33.334 0v-2.084a6.25 6.25 0 0 0-6.25-6.25h-2.084a2.083 2.083 0 0 0-2.083 2.084v10.416a2.083 2.083 0 0 0 2.083 2.084h2.084a6.25 6.25 0 0 0 6.25-6.25"/></g>`), g.Group(children),
	)
}

func Helicopter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 6.25h25zm2.083 37.5H43.75zM25 6.25v8.333zm8.333 29.167H25v8.333h8.333z"/><path stroke="#306CFE" d="M25 14.583V25a2.083 2.083 0 0 0 2.083 2.083h16.521"/><path stroke="#306CFE" d="M6.25 16.667v4.166a2.083 2.083 0 0 0 2.083 2.084h4.75a2.084 2.084 0 0 1 2.084 1.416l1.791 5.375a8.33 8.33 0 0 0 7.792 5.709h16.917a2.083 2.083 0 0 0 2.083-2.084v-6.25l-3.937-7.896a8.33 8.33 0 0 0-7.459-4.604H8.334a2.083 2.083 0 0 0-2.084 2.084"/></g>`), g.Group(children),
	)
}

func Help(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m30.02 7.458l-2.582 9m-7.458-9l2.582 9m19.98 13.562l-9-2.582m9-7.458l-9 2.582M19.98 42.542l2.582-9m7.458 9l-2.582-9M7.458 19.98l9 2.582m-9 7.458l9-2.582"/><path stroke="#306CFE" d="M43.75 25a18.75 18.75 0 1 1-37.5 0a18.75 18.75 0 0 1 37.5 0M25 16.667a8.333 8.333 0 1 0 0 16.665a8.333 8.333 0 0 0 0-16.665"/></g>`), g.Group(children),
	)
}

func HomeAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 20.833v21.459a1.604 1.604 0 0 1-1.729 1.458h-8.062V29.375h-9.584V43.75h-8.062a1.605 1.605 0 0 1-1.73-1.458V20.833"/><path stroke="#344054" d="M43.75 25L25 6.25L6.25 25"/></g>`), g.Group(children),
	)
}

func HomeAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M39.583 20.833V43.75H29.167V29.167h-8.334V43.75H10.417V20.833L25 6.25z"/>`), g.Group(children),
	)
}

func HomeTelephone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 35.417v6.25a2.083 2.083 0 0 0 2.083 2.083h14.584a2.083 2.083 0 0 0 2.083-2.083V37.5M22.917 16.667h-8.334"/><path stroke="#306CFE" d="M6.25 33.333V10.417a2.083 2.083 0 0 1 2.083-2.084H31.25v27.084H8.333a2.083 2.083 0 0 1-2.083-2.084M33.333 37.5h8.334a2.083 2.083 0 0 0 2.083-2.083V8.333a2.083 2.083 0 0 0-2.083-2.083h-8.334a2.083 2.083 0 0 0-2.083 2.083v27.084a2.083 2.083 0 0 0 2.083 2.083"/></g>`), g.Group(children),
	)
}

func HomeThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 10.417H14.583l8.334 8.333v20.833H43.75V19.604a2.08 2.08 0 0 0-.604-1.458z"/><path stroke="#344054" d="M6.25 39.583V19.604c.002-.546.22-1.07.604-1.458l7.73-7.73l8.333 8.334v20.833z"/></g>`), g.Group(children),
	)
}

func HotelTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25h-8.333v12.5H25z"/><path stroke="#306CFE" d="M43.75 43.75H6.25m27.083 0h-25V16.083A2.08 2.08 0 0 1 9.75 14l23.583-7.75zm8.334-22.917a2.083 2.083 0 0 0-2.084-2.083h-6.25v25h8.334z"/></g>`), g.Group(children),
	)
}

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 18.75a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334m6.25 14.583a6.25 6.25 0 0 0-12.5 0v2.084h12.5z"/><path stroke="#306CFE" d="M32.292 6.25h-4.167l-2.083 4.167h4.166zm-12.5 4.167h4.166L21.875 6.25h-4.167zm-7.292 0h25a2.083 2.083 0 0 1 2.083 2.083v29.167A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083V12.5a2.083 2.083 0 0 1 2.083-2.083"/></g>`), g.Group(children),
	)
}

func ImportLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 27.083h-12.5v-12.5M43.75 6.25L22.917 27.083"/><path stroke="#306CFE" d="M43.75 27.083v14.584a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h14.584"/></g>`), g.Group(children),
	)
}

func Infinite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.48 30.5a7.79 7.79 0 0 1-10.98 0L25 25l5.5-5.5a7.77 7.77 0 0 1 10.98 11m-32.96 0a7.79 7.79 0 0 0 10.98 0L25 25l-5.5-5.5a7.77 7.77 0 1 0-10.98 11"/>`), g.Group(children),
	)
}

func InfiniteTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30.917 30.875A7.77 7.77 0 1 0 30.5 19.5l-11 11a7.77 7.77 0 1 1-.417-11.375"/>`), g.Group(children),
	)
}

func InformationChatRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 16.667h-.208"/><path stroke="#344054" stroke-width="2" d="M25 33.333v-6.25"/><path stroke="#306CFE" stroke-width="2" d="M6.5 28.063a18.75 18.75 0 0 0 27.083 13.583l7.063 1.458a2.083 2.083 0 0 0 2.458-2.458l-1.437-7.084a18.75 18.75 0 1 0-35.167-5.5"/></g>`), g.Group(children),
	)
}

func InformationSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 16.667h-.208"/><path stroke="#344054" stroke-width="2" d="M25 27.083v6.25"/><path stroke="#306CFE" stroke-width="2" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Integral(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.417 39.583l.791 1.584a4.668 4.668 0 0 0 8.605-.605L25 25l5.188-15.562a4.67 4.67 0 0 1 4.437-3.188v0a4.67 4.67 0 0 1 4.167 2.583l.791 1.584"/>`), g.Group(children),
	)
}

func InvoiceDollarDoneLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.333 21.875l-3.979 5.208l-2.27-2.604m-10.417-9.896h-5.209a3.125 3.125 0 1 0 0 6.25h2.084a3.125 3.125 0 1 1 0 6.25H8.333m4.167 2.084v-2.084m8.333 8.334h12.5zM12.5 14.583V12.5z"/><path stroke="#306CFE" d="M12.5 37.5v4.167a2.083 2.083 0 0 0 2.083 2.083h25a2.083 2.083 0 0 0 2.084-2.083V8.333a2.083 2.083 0 0 0-2.084-2.083H18.75"/></g>`), g.Group(children),
	)
}

func Iron(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 35.417v6.25a2.083 2.083 0 0 1-2.084 2.083H16.667a2.083 2.083 0 0 1-2.084-2.083v-6.25z"/><path stroke="#306CFE" d="M25 4.167A28.08 28.08 0 0 1 37.5 27.52v7.896h-25V27.52A28.08 28.08 0 0 1 25 4.167"/></g>`), g.Group(children),
	)
}

func JetPlaneRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m24.333 35.27l8.48 8.48l3.729-3.73a2.08 2.08 0 0 0 .604-1.666l-1.354-14.812M19.52 9.813l-2.853 2.958m20.562 20.562l2.959-2.937zM26.563 14.208l-14.917-1.354a2.08 2.08 0 0 0-1.667.604l-3.729 3.73l8.52 8.52"/><path stroke="#306CFE" d="M19.52 36.375a2.08 2.08 0 0 0 2.96 0l16.478-16.48A16.38 16.38 0 0 0 43.75 8.334a2.083 2.083 0 0 0-2.083-2.083a16.38 16.38 0 0 0-11.563 4.792L13.625 27.52a2.08 2.08 0 0 0 0 2.958z"/></g>`), g.Group(children),
	)
}

func Jetski(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 10.417h2.833a2.08 2.08 0 0 1 1.896 1.229l3.23 7.104"/><path stroke="#306CFE" d="M27.083 39.583h-15.5a4.17 4.17 0 0 1-4.041-3.146L6.25 31.25h37.5a20.83 20.83 0 0 1-16.667 8.333m4.167-25a20.84 20.84 0 0 1-16.667 8.334h-6.25A2.083 2.083 0 0 0 6.25 25v6.25h37.5c0-12.5-12.5-16.667-12.5-16.667"/></g>`), g.Group(children),
	)
}

func Jewelry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M32.354 16.667a14.584 14.584 0 1 1-14.708 0"/><path stroke="#306CFE" d="M27.604 22.917h-5.208L14.583 12.5l5.209-6.25h10.416l5.209 6.25z"/></g>`), g.Group(children),
	)
}

func Job(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 22.917v4.166m2.083-12.5H16.667v-6.25A2.083 2.083 0 0 1 18.75 6.25h12.5a2.083 2.083 0 0 1 2.083 2.083zm2.084 8.334H14.583z"/><path stroke="#306CFE" d="M43.75 41.667v-25a2.084 2.084 0 0 0-2.083-2.084H8.333c-1.15 0-2.083.933-2.083 2.084v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083"/></g>`), g.Group(children),
	)
}

func Joystick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583V31.25zm4.167 18.75a2.083 2.083 0 0 0-2.084-2.083h-4.166a2.083 2.083 0 0 0-2.084 2.083v2.084h8.334z"/><path stroke="#306CFE" d="M39.583 43.75H10.417V37.5a2.083 2.083 0 0 1 2.083-2.083h25a2.083 2.083 0 0 1 2.083 2.083zM25 6.25a4.167 4.167 0 1 0 0 8.333a4.167 4.167 0 0 0 0-8.333"/></g>`), g.Group(children),
	)
}

func Karaoke(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 39.583H43.75M29.167 31.25H43.75zm4.166-8.333H43.75z"/><path stroke="#306CFE" d="M22.917 20.833H6.25m13.833 0l-1.208 20.959a2.083 2.083 0 0 1-2.083 1.958h-4.417a2.083 2.083 0 0 1-2.083-1.958L9.083 20.833zm-5.5-14.583a8.333 8.333 0 0 0-5.458 14.583h10.917A8.333 8.333 0 0 0 14.583 6.25"/></g>`), g.Group(children),
	)
}

func KeyEleven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 33.333h4.167M25 18.75v25zm0 22.917h4.167z"/><path stroke="#306CFE" d="M25 6.25a4.167 4.167 0 1 1 0 8.333a4.167 4.167 0 0 1 0-8.333m-8.333 12.5a4.166 4.166 0 1 0 8.332 0a4.166 4.166 0 0 0-8.332 0m12.5-4.167a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/></g>`), g.Group(children),
	)
}

func KeyLockCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 22.917a6.25 6.25 0 1 0-9.417 5.395l-.5 2.521a2.084 2.084 0 0 0 2.084 2.5h3.25a2.084 2.084 0 0 0 2.083-2.5l-.5-2.52a6.25 6.25 0 0 0 3-5.396"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func KeySeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 33.333h4.167M25 18.75v25zm0 22.917h4.167z"/><path stroke="#306CFE" d="M25 18.75a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/></g>`), g.Group(children),
	)
}

func KeyboardThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M15.52 27.083h.21m18.75 0h-.21m-9.374 0h.208"/><path stroke="#344054" stroke-width="2" d="M18.75 35.417h12.5"/><path stroke="#306CFE" stroke-width="2" d="M25 6.25v8.333m2.083 0h-4.166v4.167h4.166zM41.667 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083V20.833a2.083 2.083 0 0 1 2.083-2.083h33.334a2.083 2.083 0 0 1 2.083 2.083v20.834a2.083 2.083 0 0 1-2.083 2.083"/></g>`), g.Group(children),
	)
}

func KitchenCabinetTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 27.083h-4.166"/><path stroke="#306CFE" d="M29.167 10.417H8.333v8.333h20.834zm12.5 0H8.333v29.166h33.334zm-35.417 0h37.5zm35.417 0h-12.5v29.166h12.5z"/></g>`), g.Group(children),
	)
}

func Knife(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m21.69 31.264l-5.893-5.892l-9.56 9.56l5.892 5.893z"/><path stroke="#306CFE" d="m34.286 9.805l8.839 8.839a2.083 2.083 0 0 1 0 2.946L29.042 35.673a2.083 2.083 0 0 1-2.947 0L15.783 25.361L31.34 9.805a2.083 2.083 0 0 1 2.946 0"/></g>`), g.Group(children),
	)
}

func LaptopTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M10.417 6.25h29.166a2.083 2.083 0 0 1 2.084 2.083v18.75H8.333V8.333a2.083 2.083 0 0 1 2.084-2.083M6.25 35.417h37.5v6.25a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083z"/><path stroke="#344054" d="M41.667 27.083H8.333L6.25 35.417h37.5z"/></g>`), g.Group(children),
	)
}

func Lattern(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/><path stroke="#306CFE" d="M16.667 14.583h16.666A4.167 4.167 0 0 1 37.5 18.75v25h-25v-25a4.167 4.167 0 0 1 4.167-4.167M25 43.75v-8.333"/><path stroke="#344054" d="M8.333 43.75h33.334m-20.834-12.5A3.895 3.895 0 0 0 25 35.417a3.96 3.96 0 0 0 4.167-4.167c0-2.792-2.084-3.125-2.084-8.333c-4.166 2.083-6.25 5.396-6.25 8.333"/></g>`), g.Group(children),
	)
}

func Launch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 43.75h-8.333l1.687-6.77a4.17 4.17 0 0 1 1.73-2.46L17 32.438m16-.021l3.167 2.083a4.17 4.17 0 0 1 1.729 2.458l1.687 6.792H31.25M25 20.833v6.25"/><path stroke="#306CFE" d="M31.25 43.75h-12.5l-2.562-16.667a20.83 20.83 0 0 1 5.854-17.896l1.479-1.479a2.083 2.083 0 0 1 2.958 0l1.48 1.48a20.83 20.83 0 0 1 5.854 17.895z"/></g>`), g.Group(children),
	)
}

func LaundryBasketTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 16.667v-6.25h-12.5v6.25a2.083 2.083 0 0 0 2.083 2.083h8.334a2.083 2.083 0 0 0 2.083-2.083"/><path stroke="#306CFE" d="M6.25 10.417h37.5zm33.48 27.229l1.937-27.23H8.333l1.938 27.23a2.083 2.083 0 0 0 2.083 1.937h25.292a2.084 2.084 0 0 0 2.083-1.937"/></g>`), g.Group(children),
	)
}

func LawnMower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 22.917h11.208a2.08 2.08 0 0 1 1.875 1.145l1.5 3.021H37.5a2.083 2.083 0 0 1 2.083 2.084v4.166"/><path stroke="#306CFE" d="M6.25 8.333h2.354a2.08 2.08 0 0 1 2.084 1.792L14 33.333M35.417 37.5H18.75"/><path stroke="#344054" d="M14.583 33.333a4.166 4.166 0 1 0 0 8.333a4.166 4.166 0 0 0 0-8.333m25 0a4.166 4.166 0 1 0 0 8.333a4.166 4.166 0 0 0 0-8.333"/></g>`), g.Group(children),
	)
}

func Lcd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 33.333v7.459a55.6 55.6 0 0 1 8.334 0v-7.459z"/><path stroke="#306CFE" d="M39.583 42.708a52.2 52.2 0 0 0-29.166 0M43.75 31.25V10.417c0-1.15-.933-2.084-2.083-2.084H8.333c-1.15 0-2.083.933-2.083 2.084V31.25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083"/></g>`), g.Group(children),
	)
}

func LeftAltCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 19.604v1.23h8.333a2.083 2.083 0 0 1 2.084 2.083v4.166a2.083 2.083 0 0 1-2.084 2.084H25v1.229a2.083 2.083 0 0 1-3.562 1.458l-5.375-5.375a2.085 2.085 0 0 1 0-2.958l5.375-5.375A2.084 2.084 0 0 1 25 19.604"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func LeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 18.75L20.833 25l6.25 6.25"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func LeftDirection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.125 29.167h21.542a2.083 2.083 0 0 0 2.083-2.084v-4.166a2.084 2.084 0 0 0-2.083-2.084H20.125"/><path stroke="#306CFE" d="m20.125 20.833l6.083-7.083l-6.333-5.417L6.75 23.646a2.08 2.08 0 0 0 0 2.708l13.125 15.313l6.25-5.417l-6-7.083"/></g>`), g.Group(children),
	)
}

func LeftDirectionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 20.833h10.417v8.334H25v4.166l-9.604-6.625a2.084 2.084 0 0 1 0-3.416L25 16.667z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func LeftRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M8.333 16.667h33.334m-27.084-6.25l-6.25 6.25l6.25 6.25"/><path stroke="#306CFE" d="m35.417 39.583l6.25-6.25l-6.25-6.25m6.25 6.25H8.333"/></g>`), g.Group(children),
	)
}

func LeftRightScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 12.5v25m37.5-25v25z"/><path stroke="#306CFE" d="M31.25 20.833L35.417 25l-4.167 4.167m-12.5 0L14.583 25l4.167-4.167M35.417 25H14.583"/></g>`), g.Group(children),
	)
}

func LeftSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 43.75L29.167 25L39.583 6.25"/><path stroke="#306CFE" d="M39.583 6.25h-18.75L10.417 25l10.416 18.75h18.75"/></g>`), g.Group(children),
	)
}

func LeftSquareThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 25H31.25m-8.333-4.167L18.75 25l4.167 4.167"/><path stroke="#306CFE" d="m23.513 6.866l-16.66 16.66a2.083 2.083 0 0 0 0 2.947l16.66 16.661a2.083 2.083 0 0 0 2.946 0l16.662-16.66a2.083 2.083 0 0 0 0-2.947L26.459 6.866a2.083 2.083 0 0 0-2.946 0"/></g>`), g.Group(children),
	)
}

func Lighthouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M31.25 14.583h-12.5v8.334h12.5zm0 8.334h-12.5L16.667 43.75h16.666z"/><path stroke="#344054" d="M12.5 43.75h25m-4.167-29.167H16.667L25 6.25z"/></g>`), g.Group(children),
	)
}

func LinkAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M9.375 40.625a7.375 7.375 0 0 1 0-10.417L14.583 25A7.375 7.375 0 0 1 25 25a7.375 7.375 0 0 1 0 10.417l-5.208 5.208a7.375 7.375 0 0 1-10.417 0m27.083-16.667l5.209-5.208a7.375 7.375 0 0 0 0-10.417v0a7.375 7.375 0 0 0-10.417 0l-5.208 5.209a7.375 7.375 0 0 0 0 10.416v0a7.375 7.375 0 0 0 10.416 0"/><path stroke="#344054" d="m20.833 29.167l8.334-8.334"/></g>`), g.Group(children),
	)
}

func LinkAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M30.208 19.792a7.375 7.375 0 0 1 0 10.416L19.792 40.625a7.375 7.375 0 0 1-10.417 0a7.375 7.375 0 0 1 0-10.417"/><path stroke="#306CFE" d="M40.625 19.792a7.375 7.375 0 0 0 0-10.417a7.375 7.375 0 0 0-10.417 0L19.792 19.792a7.375 7.375 0 0 0 0 10.416v0"/></g>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 35.417v-2.084m3.125-3.125a3.125 3.125 0 1 1-6.25 0a3.125 3.125 0 0 1 6.25 0"/><path stroke="#306CFE" d="M39.583 41.667V20.833c0-1.15-.932-2.083-2.083-2.083h-25c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.932 2.083 2.083 2.083h25c1.15 0 2.083-.933 2.083-2.083m-6.25-22.917v-4.167a8.333 8.333 0 1 0-16.666 0v4.167"/></g>`), g.Group(children),
	)
}

func LockCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.75 31.25l4.167 4.167l8.333-8.334"/><path stroke="#306CFE" d="M37.5 18.75h-25c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.932 2.083 2.083 2.083h25c1.15 0 2.083-.933 2.083-2.083V20.833c0-1.15-.932-2.083-2.083-2.083m-4.167 0v-4.167a8.333 8.333 0 1 0-16.666 0v4.167"/></g>`), g.Group(children),
	)
}

func LockCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 33.333h-.208"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c5.753 0 10.417-4.664 10.417-10.417S30.753 22.917 25 22.917S14.583 27.58 14.583 33.333S19.247 43.75 25 43.75"/><path stroke="#306CFE" stroke-width="2" d="M16.667 27.083v-12.5a8.333 8.333 0 1 1 16.666 0v12.5"/></g>`), g.Group(children),
	)
}

func LockOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 18.75v-4.167a8.333 8.333 0 0 0-16.666 0v4.167"/><path stroke="#306CFE" d="M37.5 18.75h-25c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.932 2.083 2.083 2.083h25c1.15 0 2.083-.933 2.083-2.083V20.833c0-1.15-.932-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func LogIn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 8.333h6.25a2.083 2.083 0 0 1 2.083 2.084v29.166a2.083 2.083 0 0 1-2.083 2.084h-6.25"/><path stroke="#306CFE" d="m22 38.23l-3.25-2.71a2.085 2.085 0 0 1-.23-2.937l3.022-3.416H6.25v-8.334h15.292l-2.938-3.416a2.083 2.083 0 0 1 .146-2.834l3.167-2.708A2.084 2.084 0 0 1 25 12l9.98 11.646a2.084 2.084 0 0 1 0 2.708L25 38a2.083 2.083 0 0 1-3 .23"/></g>`), g.Group(children),
	)
}

func LogInDoubleArrowTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m29.167 18.75l6.25 6.25l-6.25 6.25M18.75 18.75L25 25l-6.25 6.25M25 25H6.25"/><path stroke="#306CFE" d="M35.417 41.667h6.25a2.083 2.083 0 0 0 2.083-2.084V10.417a2.083 2.083 0 0 0-2.083-2.084h-6.25"/></g>`), g.Group(children),
	)
}

func LovePlaylistTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 31.25H43.75zm0 8.333H43.75zm10.666-30.75l-.25.23l-.25-.23a1.686 1.686 0 0 0-2.395 0v0a1.71 1.71 0 0 0 0 2.417l.687.708l1.958 1.959l1.917-1.938l.688-.687a1.73 1.73 0 0 0 0-2.438a1.687 1.687 0 0 0-2.355-.02"/><path stroke="#306CFE" d="M20.833 36.458V6.25zm-7.291-7.291a7.292 7.292 0 1 0 0 14.583a7.292 7.292 0 0 0 0-14.583"/></g>`), g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 20.833V43.75m6.458-37.5l8.334 14.583zM43.75 20.833H6.25z"/><path stroke="#306CFE" d="M43.75 41.667V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083"/></g>`), g.Group(children),
	)
}

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.188 37.5c.099.31.148.633.145.958c0 2.917-3.729 5.292-8.333 5.292s-8.333-2.375-8.333-5.292a3 3 0 0 1 .145-.958"/><path stroke="#306CFE" d="M25 18.75v16.667M18.75 12.5a6.25 6.25 0 1 0 12.5 0a6.25 6.25 0 0 0-12.5 0"/></g>`), g.Group(children),
	)
}

func Maps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m31.25 12.5l-12.5-4.167L6.25 12.5v29.167l12.5-4.167l12.5 4.167l12.5-4.167V8.333z"/><path stroke="#344054" d="M31.25 41.667L18.75 37.5V8.333l12.5 4.167z"/></g>`), g.Group(children),
	)
}

func MapsLocation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M33.333 18.75h4.855a2.08 2.08 0 0 1 2.083 1.75l3.479 20.833a2.082 2.082 0 0 1-2.083 2.417H8.333a2.082 2.082 0 0 1-2.083-2.417L9.73 20.5a2.08 2.08 0 0 1 2.082-1.75h4.855"/><path stroke="#344054" d="M33.333 14.583a8.333 8.333 0 0 0-16.666 0C16.667 22.917 25 31.25 25 31.25s8.333-8.333 8.333-16.667"/></g>`), g.Group(children),
	)
}

func MapsPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 43.75h12.5"/><path stroke="#306CFE" d="M18.75 12.5A6.25 6.25 0 0 1 25 6.25v0a6.25 6.25 0 1 1-6.25 6.25M25 18.75v16.667z"/></g>`), g.Group(children),
	)
}

func Marshmallow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m25 37.5l4.167-2.083m-4.167-25V6.25zm0 16.666V43.75z"/><path stroke="#306CFE" d="M20.833 10.417h8.334A2.083 2.083 0 0 1 31.25 12.5v4.167a2.083 2.083 0 0 1-2.083 2.083h-8.334a2.083 2.083 0 0 1-2.083-2.083V12.5a2.083 2.083 0 0 1 2.083-2.083m0 8.333h8.334a2.083 2.083 0 0 1 2.083 2.083V25a2.083 2.083 0 0 1-2.083 2.083h-8.334A2.083 2.083 0 0 1 18.75 25v-4.167a2.083 2.083 0 0 1 2.083-2.083"/></g>`), g.Group(children),
	)
}

func MarshmallowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m15.396 28.708l1.479 4.417M41.063 8.938l-2.959 2.958zM8.333 41.667l18-18z"/><path stroke="#306CFE" d="m35.167 8.938l5.895 5.895a2.08 2.08 0 0 1 0 2.959l-2.958 3.041a2.084 2.084 0 0 1-2.937 0l-5.896-6a2.084 2.084 0 0 1 0-2.937l2.937-2.958a2.083 2.083 0 0 1 2.959 0m-5.896 5.895l5.896 6a2.083 2.083 0 0 1 0 2.938l-2.959 2.958a2.083 2.083 0 0 1-2.937 0l-5.896-5.896a2.084 2.084 0 0 1 0-2.937l2.958-2.959a2.084 2.084 0 0 1 2.938-.104"/></g>`), g.Group(children),
	)
}

func MaximizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m20.833 29.167l-12.5 12.5m20.834-20.834l12.5-12.5z"/><path stroke="#344054" d="M33.333 8.333h8.334v8.334m-25 25H8.333v-8.334"/></g>`), g.Group(children),
	)
}

func MeasuringTape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M41.667 29.167h-6.25V37.5h6.25a2.083 2.083 0 0 0 2.083-2.083V31.25a2.083 2.083 0 0 0-2.083-2.083M20.833 31.25a6.25 6.25 0 1 1 0-12.5a6.25 6.25 0 0 1 0 12.5"/><path stroke="#306CFE" d="M20.833 10.417A14.583 14.583 0 0 1 35.417 25v12.5a2.083 2.083 0 0 1-2.084 2.083h-12.5A14.583 14.583 0 0 1 6.25 25v0a14.583 14.583 0 0 1 14.583-14.583"/></g>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 25h37.5"/><path stroke="#306CFE" d="M6.25 12.5h25m-12.5 25h25z"/></g>`), g.Group(children),
	)
}

func MenuAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 25h37.5"/><path stroke="#306CFE" d="M6.25 12.5h37.5m-37.5 25h37.5z"/></g>`), g.Group(children),
	)
}

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M38 33.333a16.668 16.668 0 0 1-26 0m13 6.25v4.167"/><path stroke="#306CFE" d="M25 31.25a8.334 8.334 0 0 1-8.333-8.333v-8.334A8.333 8.333 0 0 1 25 6.25v0a8.333 8.333 0 0 1 8.333 8.333v8.334A8.333 8.333 0 0 1 25 31.25"/></g>`), g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M31.25 43.75h-12.5M25 35.417a8.333 8.333 0 0 1-8.333-8.334v-12.5A8.333 8.333 0 0 1 25 6.25v0a8.333 8.333 0 0 1 8.333 8.333v12.5A8.333 8.333 0 0 1 25 35.417m0 0v8.333z"/>`), g.Group(children),
	)
}

func MicrophoneAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m30.5 20.833l-1.208 20.959a2.083 2.083 0 0 1-2.084 1.958h-4.416a2.083 2.083 0 0 1-2.084-1.958L19.5 20.833zM25 6.25a8.334 8.334 0 0 0-5.458 14.583h10.916A8.334 8.334 0 0 0 25 6.25"/><path stroke="#344054" d="M33.333 20.833H16.667"/></g>`), g.Group(children),
	)
}

func MicrophoneDisable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M31.25 43.75h-12.5M25 35.417a8.334 8.334 0 0 1-8.333-8.334v-12.5A8.333 8.333 0 0 1 25 6.25v0a8.333 8.333 0 0 1 8.333 8.333v12.5A8.333 8.333 0 0 1 25 35.417m0 0v8.333z"/><path stroke="#344054" d="m8.333 8.333l33.334 33.334"/></g>`), g.Group(children),
	)
}

func MicrophoneMusicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 25a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/><path stroke="#344054" d="M39.583 14.583a6.04 6.04 0 0 0-6.25-6.25v12.5"/><path stroke="#306CFE" d="M26.354 33.333a8.33 8.33 0 0 1-7.291 1.896a9.02 9.02 0 0 1-6.563-8.458V14.938a8.604 8.604 0 0 1 6.604-8.521A8.33 8.33 0 0 1 25 7.375m4.167 36.375H12.5m8.333-8.333v8.333z"/></g>`), g.Group(children),
	)
}

func MicrophoneRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m6.25 43.75l8.542-8.542"/><path stroke="#306CFE" d="M34.896 22.917L19.208 36.875a2.083 2.083 0 0 1-2.854 0L13.23 33.75a2.083 2.083 0 0 1 0-2.854l13.854-15.792zm6.416-14.209a8.33 8.33 0 0 0-11.791 0a8.34 8.34 0 0 0-2.438 6.438l7.709 7.708a8.33 8.33 0 0 0 6.458-2.375a8.334 8.334 0 0 0 .063-11.792z"/><path stroke="#344054" d="M36.896 24.896L25.104 13.104"/></g>`), g.Group(children),
	)
}

func MicrophoneStand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 43.75h-8.334M25 27.083V43.75z"/><path stroke="#306CFE" d="M31.833 22.188L13.792 32.875A2.083 2.083 0 0 1 11 32.25l-2.48-3.646A2.083 2.083 0 0 1 9 25.792l16.667-12.75zm8.563-12.271a8.333 8.333 0 0 0-11.563-2.25a8.33 8.33 0 0 0-3.416 4.979l6.833 10.146a8.33 8.33 0 0 0 5.896-1.292a8.333 8.333 0 0 0 2.25-11.583"/><path stroke="#344054" d="m24.104 10.708l9.313 13.813"/></g>`), g.Group(children),
	)
}

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m41.667 8.333l-12.5 12.5M8.333 41.667l12.5-12.5z"/><path stroke="#344054" d="M12.5 29.167h8.333V37.5M37.5 20.833h-8.333V12.5"/></g>`), g.Group(children),
	)
}

func MinimizeLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m8.333 8.333l12.5 12.5m20.834 20.834l-12.5-12.5z"/><path stroke="#344054" d="M29.167 37.5v-8.333H37.5M20.833 12.5v8.333H12.5"/></g>`), g.Group(children),
	)
}

func MinimizeSize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 6.25L31.854 18.146M18.146 31.854L6.25 43.75zm0-13.708L6.25 6.25zM43.75 43.75L31.854 31.854z"/><path stroke="#306CFE" d="M8.333 31.25h8.334a2.083 2.083 0 0 1 2.083 2.083v8.334m0-33.334v8.334a2.083 2.083 0 0 1-2.083 2.083H8.333M31.25 41.667v-8.334a2.083 2.083 0 0 1 2.083-2.083h8.334m0-12.5h-8.334a2.083 2.083 0 0 1-2.083-2.083V8.333"/></g>`), g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M39.583 25H10.417"/>`), g.Group(children),
	)
}

func MinusCollection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 20.833h12.5"/><path stroke="#306CFE" d="M14.583 43.75h27.084a2.083 2.083 0 0 0 2.083-2.083v-31.25"/><path stroke="#306CFE" d="M33.333 6.25h-25c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.084 2.083 2.084h25c1.15 0 2.084-.933 2.084-2.084v-25c0-1.15-.933-2.083-2.084-2.083"/></g>`), g.Group(children),
	)
}

func MobileDataCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m14.583 27.083l4.167 4.167l4.167-4.167m12.5-6.25l-4.167-4.166l-4.167 4.166M31.25 31.25V16.667m-12.5 0V31.25z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func MobileHotspot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="M18.75 43.75h12.5M25 27.083V43.75z"/><path stroke="#344054" stroke-width="3" d="M24.896 25h.208"/><path stroke="#344054" stroke-width="2" d="M11.75 38.25a18.75 18.75 0 1 1 26.5 0"/><path stroke="#344054" stroke-width="2" d="M32.375 32.375a10.416 10.416 0 1 0-14.75 0"/></g>`), g.Group(children),
	)
}

func MobilePaymentDollar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 22.917h-5.208a3.125 3.125 0 0 0-3.125 3.125v0a3.125 3.125 0 0 0 3.125 3.125h2.083a3.125 3.125 0 0 1 3.125 3.125v0a3.125 3.125 0 0 1-3.125 3.125H31.25m4.167 2.083v-2.083m0-12.5v-2.084z"/><path stroke="#306CFE" d="M29.167 43.75H12.5a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 12.5 6.25h20.833a2.083 2.083 0 0 1 2.084 2.083V12.5"/><path stroke="#306CFE" d="M27.083 6.25H18.75l1.042 4.167h6.25z"/></g>`), g.Group(children),
	)
}

func MobilePaymentDoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 22.917h12.5M8.333 14.583H18.75v20.834H8.333a2.083 2.083 0 0 1-2.083-2.084V16.667a2.084 2.084 0 0 1 2.083-2.084M27.083 25l2.271 2.604l3.98-5.208"/><path stroke="#306CFE" d="M43.75 41.667V8.333a2.083 2.083 0 0 0-2.083-2.083H20.833a2.083 2.083 0 0 0-2.083 2.083v33.334a2.083 2.083 0 0 0 2.083 2.083h20.834a2.083 2.083 0 0 0 2.083-2.083m-15.625-31.25h6.25l1.042-4.167h-8.334z"/></g>`), g.Group(children),
	)
}

func Money(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 12.5a2.083 2.083 0 0 0-2.083-2.083h-6.25a8.333 8.333 0 0 0 8.333 8.333zm-37.5 6.25a8.333 8.333 0 0 0 8.333-8.333h-6.25A2.083 2.083 0 0 0 6.25 12.5zm37.5 12.5a8.334 8.334 0 0 0-8.333 8.333h6.25A2.083 2.083 0 0 0 43.75 37.5zM6.25 37.5a2.083 2.083 0 0 0 2.083 2.083h6.25A8.333 8.333 0 0 0 6.25 31.25zM25 20.833a4.167 4.167 0 1 0 0 8.334a4.167 4.167 0 0 0 0-8.334"/><path stroke="#306CFE" d="M41.667 10.417H8.333c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func MoneyAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M20.73 20.833h.207"/><path stroke="#306CFE" stroke-width="2" d="M33.333 10.417h-25c-1.15 0-2.083.933-2.083 2.083v16.667c0 1.15.933 2.083 2.083 2.083h25c1.15 0 2.084-.933 2.084-2.083V12.5c0-1.15-.933-2.083-2.084-2.083"/><path stroke="#306CFE" stroke-width="2" d="M35.417 18.75h6.25a2.083 2.083 0 0 1 2.083 2.083V37.5a2.083 2.083 0 0 1-2.083 2.083h-25a2.083 2.083 0 0 1-2.084-2.083v-6.25"/></g>`), g.Group(children),
	)
}

func MoneyAltOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 14.583h-5.209a3.125 3.125 0 0 0-3.125 3.125v0a3.125 3.125 0 0 0 3.125 3.125h2.084a3.125 3.125 0 0 1 3.125 3.125v0a3.125 3.125 0 0 1-3.125 3.125h-5.209m4.167 0v2.084M25 12.5v2.083z"/><path stroke="#306CFE" d="M12.5 43.75h25m4.167-37.5H8.333c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.084 2.083 2.084h33.334c1.15 0 2.083-.933 2.083-2.084v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func MoneyBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.667 14.583s5.979-.479 8.5 1.188a7.38 7.38 0 0 0 5.5.896"/><path stroke="#306CFE" d="m27.667 14.583l1.062-5.875a2.082 2.082 0 0 0-2.062-2.458h-7.5a2.083 2.083 0 0 0-2.084 2.458l1.084 5.875zm-4.75 29.167a30 30 0 0 0 8.729-1.187a8.33 8.33 0 0 0 5.854-7.959a14.3 14.3 0 0 0-1.5-6.25l-6.25-12.5a2.08 2.08 0 0 0-1.875-1.146h-9.917a2.08 2.08 0 0 0-1.875 1.146l-6.25 12.5a14.3 14.3 0 0 0-1.5 6.25v0a8.33 8.33 0 0 0 5.855 7.959a30 30 0 0 0 8.729 1.187"/></g>`), g.Group(children),
	)
}

func MoneyDollar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 18.75h-5.209a3.125 3.125 0 0 0-3.125 3.125v0A3.125 3.125 0 0 0 23.958 25h2.084a3.125 3.125 0 0 1 3.125 3.125v0a3.125 3.125 0 0 1-3.125 3.125h-5.209M25 33.333V31.25M43.75 12.5a2.083 2.083 0 0 0-2.083-2.083h-6.25a8.333 8.333 0 0 0 8.333 8.333zm-37.5 6.25a8.333 8.333 0 0 0 8.333-8.333h-6.25A2.083 2.083 0 0 0 6.25 12.5zm37.5 12.5a8.333 8.333 0 0 0-8.333 8.333h6.25A2.083 2.083 0 0 0 43.75 37.5zM6.25 37.5a2.083 2.083 0 0 0 2.083 2.083h6.25A8.333 8.333 0 0 0 6.25 31.25zM25 18.75v-2.083z"/><path stroke="#306CFE" d="M41.667 10.417H8.333c-1.15 0-2.083.932-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func MoneyDollarCoin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 31.25h5.209a3.125 3.125 0 0 0 3.125-3.125v0A3.125 3.125 0 0 0 26.042 25h-2.084a3.125 3.125 0 1 1 0-6.25h5.209M25 33.333V31.25m0-12.5v-2.083z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func MoneyEuro(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 20.354A6.25 6.25 0 0 0 25 18.75a6.25 6.25 0 0 0 0 12.5a6.25 6.25 0 0 0 4.167-1.604M16.667 25h6.25zM43.75 12.5a2.083 2.083 0 0 0-2.083-2.083h-6.25a8.333 8.333 0 0 0 8.333 8.333zm-37.5 6.25a8.333 8.333 0 0 0 8.333-8.333h-6.25A2.083 2.083 0 0 0 6.25 12.5zm37.5 12.5a8.334 8.334 0 0 0-8.333 8.333h6.25A2.083 2.083 0 0 0 43.75 37.5zM6.25 37.5a2.083 2.083 0 0 0 2.083 2.083h6.25A8.333 8.333 0 0 0 6.25 31.25z"/><path stroke="#306CFE" d="M41.667 10.417H8.333c-1.15 0-2.083.933-2.083 2.083v25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083v-25c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func MoneyThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M20.73 20.833h.207"/><path stroke="#306CFE" stroke-width="2" d="M43.75 18.75V37.5a2.083 2.083 0 0 1-2.083 2.083h-31.25"/><path stroke="#306CFE" stroke-width="2" d="M8.333 31.25h25c1.15 0 2.084-.933 2.084-2.083V12.5c0-1.15-.933-2.083-2.084-2.083h-25c-1.15 0-2.083.932-2.083 2.083v16.667c0 1.15.933 2.083 2.083 2.083"/></g>`), g.Group(children),
	)
}

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.333 41.667l-2.083-8.334h-12.5l-2.083 8.334z"/><path stroke="#306CFE" d="M12.5 41.667h25zm29.167-8.334H8.333A2.083 2.083 0 0 1 6.25 31.25V10.417a2.083 2.083 0 0 1 2.083-2.084h33.334a2.083 2.083 0 0 1 2.083 2.084V31.25a2.083 2.083 0 0 1-2.083 2.083"/></g>`), g.Group(children),
	)
}

func MoreCircleVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25 14.48v.207m0 10.209v.208m0 10.209v.208"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#344054" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M25 12.396v.208m0 12.292v.208m0 12.292v.208"/>`), g.Group(children),
	)
}

func Mosque(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M28.125 11.438a3.675 3.675 0 1 1-5.208-5.188M25 12.5v6.25"/><path stroke="#306CFE" d="M39.583 43.75a18.6 18.6 0 0 0 2.084-7.646c0-4.166-1.917-6.77-6.75-9.396A36.3 36.3 0 0 1 25 18.75a36.3 36.3 0 0 1-9.917 7.958c-4.833 2.625-6.75 5.313-6.75 9.396a18.6 18.6 0 0 0 2.084 7.646z"/></g>`), g.Group(children),
	)
}

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.5 18.75a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/><path stroke="#306CFE" d="M8.333 43.75h14.584l4.937-10.125L22.667 24A2.083 2.083 0 0 0 19 24zm30 0H22.917l9.375-19.23l7.916 16.23a2.083 2.083 0 0 1-1.875 3"/><path stroke="#344054" d="M6.25 43.75h37.5"/></g>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583v2.084"/><path stroke="#306CFE" d="M25 6.25a13.79 13.79 0 0 1 13.77 14.583l-.583 10.417A13.187 13.187 0 0 1 25 43.75v0a13.187 13.187 0 0 1-13.187-12.5l-.584-10.417A13.79 13.79 0 0 1 25 6.25"/></g>`), g.Group(children),
	)
}

func MpThreePlayer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m35.417 27.083l8.333 2.084V25l-8.333-2.083v16.666zm-4.167 8.334a4.167 4.167 0 1 0 0 8.333a4.167 4.167 0 0 0 0-8.333"/><path stroke="#306CFE" d="M18.75 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h25a2.083 2.083 0 0 1 2.084 2.083v6.25"/><path stroke="#306CFE" d="M20.833 27.083a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/></g>`), g.Group(children),
	)
}

func MultiFolder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M14.583 35.417h-6.25a2.083 2.083 0 0 1-2.083-2.084v-25A2.083 2.083 0 0 1 8.333 6.25h8.334a2.08 2.08 0 0 1 1.479.604l5.02 5.042a2.08 2.08 0 0 0 1.48.604h10.77a2.083 2.083 0 0 1 2.084 2.083v6.25"/><path stroke="#344054" d="M41.667 20.833h-25c-1.15 0-2.084.933-2.084 2.084v18.75c0 1.15.933 2.083 2.084 2.083h25c1.15 0 2.083-.933 2.083-2.083v-18.75c0-1.15-.933-2.084-2.083-2.084"/></g>`), g.Group(children),
	)
}

func MusicAlbumTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 33.333a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/><path stroke="#344054" d="M22.917 29.167v-12.5a6.04 6.04 0 0 1 6.25 6.25"/><path stroke="#306CFE" d="M37.5 41.667H8.333a2.083 2.083 0 0 1-2.083-2.084V10.417a2.083 2.083 0 0 1 2.083-2.084H37.5a2.083 2.083 0 0 1 2.083 2.084v29.166a2.083 2.083 0 0 1-2.083 2.084m2.083-25v16.666a10.417 10.417 0 0 0 0-16.666"/></g>`), g.Group(children),
	)
}

func MusicDisable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M27.083 36.458V6.25a54 54 0 0 0 7.167 6.25c3.458 2.23 3.98 5.75 2.396 10.417"/><path stroke="#306CFE" d="M19.792 43.75a7.292 7.292 0 1 0 0-14.583a7.292 7.292 0 0 0 0 14.583"/><path stroke="#344054" d="M39.583 43.75L10.417 6.25"/></g>`), g.Group(children),
	)
}

func MusicFileTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 43.75a4.167 4.167 0 1 0 0-8.333a4.167 4.167 0 0 0 0 8.333"/><path stroke="#344054" d="M41.667 33.333a6.04 6.04 0 0 0-6.25-6.25v12.5"/><path stroke="#306CFE" d="M18.75 39.583h-8.333A2.083 2.083 0 0 1 8.333 37.5V18.75l12.5-12.5h12.5a2.083 2.083 0 0 1 2.084 2.083V18.75"/><path stroke="#306CFE" d="M20.833 18.75V6.25l-12.5 12.5z"/></g>`), g.Group(children),
	)
}

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M27.083 36.458V6.25a54 54 0 0 0 7.167 6.25c3.458 2.23 3.98 5.75 2.396 10.417"/><path d="M19.792 43.75a7.292 7.292 0 1 0 0-14.583a7.292 7.292 0 0 0 0 14.583"/></g>`), g.Group(children),
	)
}

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m35.417 6.25l-14.25 10.417H12.5a2.083 2.083 0 0 0-2.083 2.083v12.5a2.083 2.083 0 0 0 2.083 2.083h8.667l14.25 10.417z"/><path stroke="#344054" d="M43.75 38.813L6.25 11.186"/></g>`), g.Group(children),
	)
}

func Next(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m35.417 25l-25 18.75V6.25z"/>`), g.Group(children),
	)
}

func NoteAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 31.25h12.5M29.167 12.5h-8.334a2.083 2.083 0 0 1-2.083-2.083V8.333a2.083 2.083 0 0 1 2.083-2.083h8.334a2.083 2.083 0 0 1 2.083 2.083v2.084a2.083 2.083 0 0 1-2.083 2.083M18.75 22.917h12.5z"/><path stroke="#306CFE" d="M37.5 8.333a2.083 2.083 0 0 1 2.083 2.084v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083v-31.25A2.083 2.083 0 0 1 12.5 8.333"/></g>`), g.Group(children),
	)
}

func NoteBook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 6.25h-31.25c-1.15 0-2.084.933-2.084 2.083v33.334c0 1.15.933 2.083 2.084 2.083h31.25c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/><path stroke="#344054" d="M10.417 16.667H6.25zm0 8.333H6.25zm0 8.333H6.25zm25-18.75H18.75v8.334h16.667z"/></g>`), g.Group(children),
	)
}

func NoteDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 33.333v-12.5m4.167-8.333h-8.334a2.083 2.083 0 0 1-2.083-2.083V8.333a2.083 2.083 0 0 1 2.083-2.083h8.334a2.083 2.083 0 0 1 2.083 2.083v2.084a2.083 2.083 0 0 1-2.083 2.083m-8.334 16.667L25 33.333l4.167-4.166"/><path stroke="#306CFE" d="M37.5 8.333a2.083 2.083 0 0 1 2.083 2.084v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083v-31.25A2.083 2.083 0 0 1 12.5 8.333"/></g>`), g.Group(children),
	)
}

func NoteText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 20.833v12.5M29.167 12.5h-8.334a2.083 2.083 0 0 1-2.083-2.083V8.333a2.083 2.083 0 0 1 2.083-2.083h8.334a2.083 2.083 0 0 1 2.083 2.083v2.084a2.083 2.083 0 0 1-2.083 2.083M18.75 20.833h12.5z"/><path stroke="#306CFE" d="M37.5 8.333a2.083 2.083 0 0 1 2.083 2.084v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083v-31.25A2.083 2.083 0 0 1 12.5 8.333"/></g>`), g.Group(children),
	)
}

func NoteUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 20.833v12.5M29.167 12.5h-8.334a2.083 2.083 0 0 1-2.083-2.083V8.333a2.083 2.083 0 0 1 2.083-2.083h8.334a2.083 2.083 0 0 1 2.083 2.083v2.084a2.083 2.083 0 0 1-2.083 2.083M20.833 25L25 20.833L29.167 25"/><path stroke="#306CFE" d="M37.5 8.333a2.083 2.083 0 0 1 2.083 2.084v31.25A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083v-31.25A2.083 2.083 0 0 1 12.5 8.333"/></g>`), g.Group(children),
	)
}

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 6.25H10.417c-1.15 0-2.084.933-2.084 2.083v33.334c0 1.15.933 2.083 2.084 2.083h29.166c1.15 0 2.084-.933 2.084-2.083V8.333c0-1.15-.933-2.083-2.084-2.083"/><path stroke="#344054" d="M10.417 6.25h6.25v37.5h-6.25a2.083 2.083 0 0 1-2.084-2.083V8.333a2.083 2.083 0 0 1 2.084-2.083"/></g>`), g.Group(children),
	)
}

func NotificationAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M38.542 16.667a5.208 5.208 0 1 0 0-10.417a5.208 5.208 0 0 0 0 10.417"/><path stroke="#306CFE" d="M41.667 25v12.5a4.167 4.167 0 0 1-4.167 4.167h-25A4.167 4.167 0 0 1 8.333 37.5v-25A4.167 4.167 0 0 1 12.5 8.333H25"/></g>`), g.Group(children),
	)
}

func NotificationBell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 37.5h-12.5a6.25 6.25 0 0 0 12.5 0"/><path stroke="#306CFE" d="M40.375 29.958a4.42 4.42 0 0 1 1.292 3.125v0A4.417 4.417 0 0 1 37.25 37.5h-24.5a4.416 4.416 0 0 1-4.417-4.417v0a4.42 4.42 0 0 1 1.292-3.125l2.875-2.875V18.75A12.5 12.5 0 0 1 25 6.25v0a12.5 12.5 0 0 1 12.5 12.5v8.333z"/></g>`), g.Group(children),
	)
}

func NotificationCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 37.5h-12.5a6.25 6.25 0 0 0 12.5 0"/><path stroke="#306CFE" d="M37.5 18.75v8.333l2.875 2.875A4.417 4.417 0 0 1 37.25 37.5h-24.5a4.416 4.416 0 0 1-3.125-7.542l2.875-2.875V18.75A12.5 12.5 0 0 1 25 6.25a12.5 12.5 0 0 1 6 1.52"/><path stroke="#344054" d="M29.167 12.5a6.25 6.25 0 1 0 12.5 0a6.25 6.25 0 0 0-12.5 0"/></g>`), g.Group(children),
	)
}

func OneStPlace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 18.75L25 16.667v16.666m-2.083 0h4.166"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Package(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 16.667h-12.5V31.25L25 29.167l6.25 2.083z"/><path stroke="#306CFE" d="M43.75 16.667v25a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083v-25m0 0l9.813-9.813a2.08 2.08 0 0 1 1.458-.604h14.958c.547.002 1.07.22 1.459.604l9.812 9.813z"/></g>`), g.Group(children),
	)
}

func PackageAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M6.25 16.667h37.5v25a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083z"/><path stroke="#344054" d="M25 6.25v20.833"/><path stroke="#306CFE" d="m43.75 16.667l-9.812-9.813a2.08 2.08 0 0 0-1.459-.604H17.521c-.547.002-1.07.22-1.459.604L6.25 16.667"/></g>`), g.Group(children),
	)
}

func PaintRoller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 10.417H8.333A2.083 2.083 0 0 0 6.25 12.5v6.458a2.083 2.083 0 0 0 1.875 2.084l17.083 1.708a2.083 2.083 0 0 1 1.875 2.083v4.334"/><path stroke="#306CFE" d="M12.5 12.5V8.333a2.083 2.083 0 0 1 2.083-2.083h25a2.083 2.083 0 0 1 2.084 2.083V12.5a2.083 2.083 0 0 1-2.084 2.083h-25A2.083 2.083 0 0 1 12.5 12.5m14.583 31.25a4.167 4.167 0 0 0 4.167-4.167V31.25a2.083 2.083 0 0 0-2.083-2.083H25a2.083 2.083 0 0 0-2.083 2.083v8.333a4.167 4.167 0 0 0 4.166 4.167"/></g>`), g.Group(children),
	)
}

func Pantone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M6.25 37.5V8.333A2.083 2.083 0 0 1 8.333 6.25h8.334a2.083 2.083 0 0 1 2.083 2.083v14.084"/><path stroke="#344054" d="m9.917 31.25l18.75-18.75a2.083 2.083 0 0 1 2.937 0l5.896 5.854a2.083 2.083 0 0 1 0 2.938l-9.917 9.958"/><path stroke="#306CFE" d="M6.25 37.5a6.25 6.25 0 0 1 6.25-6.25h29.167a2.083 2.083 0 0 1 2.083 2.083v8.334a2.083 2.083 0 0 1-2.083 2.083H12.5a6.25 6.25 0 0 1-6.25-6.25"/></g>`), g.Group(children),
	)
}

func ParabolicFunction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 6.25c.458 14.27 2.541 25 6.25 25s5.791-10.73 6.25-25"/><path stroke="#306CFE" d="M43.75 39.583H6.25M10.417 6.25v37.5z"/></g>`), g.Group(children),
	)
}

func ParkingSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 35.417V14.583H25a6.25 6.25 0 1 1 0 12.5h-6.25"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.583 6.25h4.167a2.083 2.083 0 0 1 2.083 2.083v33.334a2.083 2.083 0 0 1-2.083 2.083h-4.167a2.083 2.083 0 0 1-2.083-2.083V8.333a2.083 2.083 0 0 1 2.083-2.083m16.667 37.5h4.167a2.083 2.083 0 0 0 2.083-2.083V8.333a2.083 2.083 0 0 0-2.083-2.083H31.25a2.083 2.083 0 0 0-2.083 2.083v33.334a2.083 2.083 0 0 0 2.083 2.083"/>`), g.Group(children),
	)
}

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 16.667v16.666m8.334-16.666v16.666z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func PencilRuler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m43.146 9.77l-2.917-2.916a2.083 2.083 0 0 0-2.916 0l-10.23 10.23v5.833h5.834l10.229-10.23a2.083 2.083 0 0 0 0-2.916"/><path stroke="#306CFE" d="M14.583 6.25v37.5H6.25V6.25zm25 29.167h-25v8.333h25z"/></g>`), g.Group(children),
	)
}

func Pendulum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 6.25v29.167m0 8.333a4.167 4.167 0 1 1 0-8.334a4.167 4.167 0 0 1 0 8.334"/><path stroke="#344054" d="M10.417 6.25h29.166"/></g>`), g.Group(children),
	)
}

func PendulumFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M15.833 35.604L25 6.25m10.417 33.333a4.166 4.166 0 1 0-8.332 0a4.166 4.166 0 0 0 8.332 0M31.25 6.25v29.167zm-16.667 37.5a4.167 4.167 0 1 0 0-8.334a4.167 4.167 0 0 0 0 8.334"/><path stroke="#344054" d="M10.417 6.25h29.166"/></g>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.75 31.25v8.187a4.166 4.166 0 0 1-4.77 4.167A37.5 37.5 0 0 1 6.541 11.021a4.167 4.167 0 0 1 4.146-4.771h8.062a2.083 2.083 0 0 1 2.083 1.854c.207 2.73.913 5.4 2.084 7.875a2.083 2.083 0 0 1-.875 2.625l-1.792 1.02a2.085 2.085 0 0 0-.687 3.043a29.3 29.3 0 0 0 7.687 7.687a2.083 2.083 0 0 0 3.042-.687l1.02-1.792a2.083 2.083 0 0 1 2.709-.792a22.4 22.4 0 0 0 7.875 2.084a2.083 2.083 0 0 1 1.854 2.083"/>`), g.Group(children),
	)
}

func PhoneSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m34.042 34.438l.895-1.376a2.897 2.897 0 0 0-1.062-4.166a26 26 0 0 1-1.5-.854a2.917 2.917 0 0 0-3.917.75l-2.666 3.791a19.6 19.6 0 0 0 6.25 2.792a1.71 1.71 0 0 0 2-.937M14.583 17.854a19.6 19.6 0 0 0 2.792 6.25l3.792-2.666a2.916 2.916 0 0 0 .75-3.917c-.292-.48-.584-.98-.855-1.5a2.895 2.895 0 0 0-4.166-1.063l-1.375.896a1.71 1.71 0 0 0-.938 2m11.209 14.729a28 28 0 0 1-4.625-3.75a28 28 0 0 1-3.75-4.625"/><path stroke="#306CFE" d="M39.583 6.25H10.417a4.167 4.167 0 0 0-4.167 4.167v29.166a4.167 4.167 0 0 0 4.167 4.167h29.166a4.167 4.167 0 0 0 4.167-4.167V10.417a4.167 4.167 0 0 0-4.167-4.167"/></g>`), g.Group(children),
	)
}

func PicnicBasketThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M32.042 43.75H17.958a8.333 8.333 0 0 1-8.333-7.5L8.333 22.917h33.334L40.333 36.25a8.334 8.334 0 0 1-8.291 7.5M25 6.25a14.583 14.583 0 0 0-14.583 14.583v2.084h29.166v-2.084A14.583 14.583 0 0 0 25 6.25"/><path stroke="#344054" d="M6.25 22.917h37.5"/></g>`), g.Group(children),
	)
}

func PicnicBasketTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M32.042 43.75H17.958a8.333 8.333 0 0 1-8.333-7.5L8.333 22.917h33.334L40.333 36.25a8.334 8.334 0 0 1-8.291 7.5M25 6.25a14.583 14.583 0 0 0-14.583 14.583v2.084h29.166v-2.084A14.583 14.583 0 0 0 25 6.25m1.73 26.563l6.603-9.896H16.667l6.604 9.895a2.08 2.08 0 0 0 3.458 0"/><path stroke="#344054" d="M6.25 22.917h37.5"/></g>`), g.Group(children),
	)
}

func PipeFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M8.333 33.333v-25h8.334v25H25V22.917a8.333 8.333 0 0 1 8.333-8.334H43.75v8.334H33.333v10.416A8.333 8.333 0 0 1 25 41.667h-8.333a8.333 8.333 0 0 1-8.334-8.334"/><path stroke="#344054" d="M18.75 8.333H6.25M43.75 25V12.5z"/></g>`), g.Group(children),
	)
}

func PlaceDisable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 43.75s-12.5-14.583-12.5-25a12.5 12.5 0 0 1 25 0c0 10.417-12.5 25-12.5 25"/><path stroke="#344054" d="M8.333 8.333L41.667 37.5"/></g>`), g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.333 25l-12.5 8.333V16.667z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Plunger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25a10.3 10.3 0 0 1 4.167.875V10.417a4.167 4.167 0 1 0-8.334 0v21.708A10.3 10.3 0 0 1 25 31.25"/><path stroke="#306CFE" d="M12.5 43.75h25m-2.083 0H14.583v-2.083a10.417 10.417 0 0 1 20.834 0z"/></g>`), g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25 10.417v29.166M10.417 25h29.166z"/>`), g.Group(children),
	)
}

func Podium(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m24.042 8.208l-2.167.313l1.563 1.52l-.375 2.146l1.937-1l1.938 1l-.375-2.145l1.562-1.521l-2.167-.313L25 6.25z"/><path stroke="#306CFE" d="M18.75 43.75H6.25V33.333h12.5zm12.5-22.917h-12.5V43.75h12.5zm12.5 8.334h-12.5V43.75h12.5z"/></g>`), g.Group(children),
	)
}

func Power(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 6.25V25"/><path stroke="#306CFE" d="M36.792 15.292a16.667 16.667 0 1 1-23.584 0"/></g>`), g.Group(children),
	)
}

func PowerBank(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.75 25l2.083-4.167h-4.166l2.083-4.166m0 18.75v6.25a2.083 2.083 0 0 0 2.083 2.083H37.5a2.083 2.083 0 0 0 2.083-2.083V18.75"/><path stroke="#306CFE" d="M29.167 33.333v-25a2.083 2.083 0 0 0-2.084-2.083H10.417a2.083 2.083 0 0 0-2.084 2.083v25a2.083 2.083 0 0 0 2.084 2.084h16.666a2.083 2.083 0 0 0 2.084-2.084m12.5-22.916H37.5v6.25h4.167z"/></g>`), g.Group(children),
	)
}

func PresentCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 27.083a5.208 5.208 0 1 0 0-10.416a5.208 5.208 0 0 0 0 10.416"/><path stroke="#306CFE" d="m28.333 35.417l5 6.25m6.25-6.25H10.417a2.083 2.083 0 0 1-2.084-2.084v-25h33.334v25a2.083 2.083 0 0 1-2.084 2.084M6.25 8.333h37.5zm15.417 27.084l-5 6.25z"/></g>`), g.Group(children),
	)
}

func PresentGrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 27.083c8.333 0 11.937-5.791 14.583-10.416"/><path stroke="#344054" d="M25 16.667h6.25l2.083 6.25"/><path stroke="#306CFE" d="m28.333 35.417l5 6.25m6.25-6.25H10.417a2.083 2.083 0 0 1-2.084-2.084v-25h33.334v25a2.083 2.083 0 0 1-2.084 2.084M6.25 8.333h37.5zm15.417 27.084l-5 6.25z"/></g>`), g.Group(children),
	)
}

func PresentationBarChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 25v-6.25M25 33.333V43.75zm2.083 10.417h-4.166zM25 25V14.583zm8.333 0v-6.25z"/><path stroke="#306CFE" d="M6.25 6.25h37.5m-35.417 0h33.334v25a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083z"/></g>`), g.Group(children),
	)
}

func PriceTag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M34.27 15.625h.21"/><path stroke="#306CFE" stroke-width="2" d="M43.75 20.833v-12.5a2.083 2.083 0 0 0-2.083-2.083h-12.5l-21.5 21.375a2.083 2.083 0 0 0 0 2.938l11.791 11.791a2.084 2.084 0 0 0 2.938 0z"/></g>`), g.Group(children),
	)
}

func PrintAltNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 33.333H8.333A2.083 2.083 0 0 1 6.25 31.25V14.583A2.083 2.083 0 0 1 8.333 12.5h33.334a2.083 2.083 0 0 1 2.083 2.083V31.25a2.083 2.083 0 0 1-2.083 2.083H37.5M16.667 6.25h16.666"/><path stroke="#344054" d="M37.5 20.833h-25V43.75h25z"/><path stroke="#344054" d="m20.833 33.333l2.084 2.084l6.25-6.25"/></g>`), g.Group(children),
	)
}

func PrintAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.604 33.333h6.063a2.083 2.083 0 0 0 2.083-2.083V14.583a2.083 2.083 0 0 0-2.083-2.083H8.333a2.083 2.083 0 0 0-2.083 2.083V31.25a2.083 2.083 0 0 0 2.083 2.083h6.063M35.417 6.25H14.583v6.25h20.834z"/><path stroke="#344054" d="M37.5 43.75h-25l4.167-22.917h16.666z"/></g>`), g.Group(children),
	)
}

func Processor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.5 33.333H8.333M25 12.5V6.25zm8.333 0V8.333zm-16.666 0V8.333zM37.5 25h6.25zm0 8.333h4.167zm0-16.666h4.167zM25 37.5v6.25zm-8.333 0v4.167zm16.666 0v4.167zM12.5 25H6.25zm0-8.333H8.333z"/><path stroke="#306CFE" d="M35.417 37.5H14.583a2.083 2.083 0 0 1-2.083-2.083V14.583a2.083 2.083 0 0 1 2.083-2.083h20.834a2.083 2.083 0 0 1 2.083 2.083v20.834a2.083 2.083 0 0 1-2.083 2.083m-6.25-16.667h-8.334v8.334h8.334z"/></g>`), g.Group(children),
	)
}

func PushPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25v12.5"/><path stroke="#306CFE" d="M14.583 6.25h20.834m-.23 17.604a4.17 4.17 0 0 1 2.313 3.73v3.666h-25v-3.667a4.17 4.17 0 0 1 2.313-3.729l1.854-.937L18.75 6.25h12.5l2.083 16.667z"/></g>`), g.Group(children),
	)
}

func QrCodeScanTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 22.917h-8.334v-8.334h8.334zm12.5-8.334H31.25v4.167h4.167zm-8.334 20.834h8.334v-8.334h-8.334zm-12.5 0h4.167V31.25h-4.167z"/><path stroke="#306CFE" d="M16.667 6.25H8.333A2.083 2.083 0 0 0 6.25 8.333v8.334m37.5 0V8.333a2.083 2.083 0 0 0-2.083-2.083h-8.334M6.25 33.333v8.334a2.083 2.083 0 0 0 2.083 2.083h8.334m16.666 0h8.334a2.083 2.083 0 0 0 2.083-2.083v-8.334"/></g>`), g.Group(children),
	)
}

func Question(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M25 27.083c0-4.166 6.25-2.083 6.25-8.333s-12.5-6.25-12.5 0"/><path stroke="#344054" stroke-width="3" d="M25.104 35.417h-.208"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Radius(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 25h-.208"/><path stroke="#344054" stroke-width="2" d="M43.75 25H25"/><path stroke="#306CFE" stroke-width="2" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func ReceiptAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M32.292 19.792H21.875m5.208-5.209V25z"/><path stroke="#306CFE" d="M43.75 8.333v30.209a5.208 5.208 0 1 1-10.417 0v-5.209H10.417v-25A2.083 2.083 0 0 1 12.5 6.25h29.167a2.083 2.083 0 0 1 2.083 2.083M33.333 38.542v-5.209H6.25v5.209a5.21 5.21 0 0 0 5.208 5.208h27.084a5.21 5.21 0 0 1-5.209-5.208"/></g>`), g.Group(children),
	)
}

func ReceiptCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 20.833L25 25l8.333-8.333"/><path stroke="#306CFE" d="M43.75 8.333v30.209a5.208 5.208 0 1 1-10.417 0v-5.209H10.417v-25A2.083 2.083 0 0 1 12.5 6.25h29.167a2.083 2.083 0 0 1 2.083 2.083M33.333 38.542v-5.209H6.25v5.209a5.21 5.21 0 0 0 5.208 5.208h27.084a5.21 5.21 0 0 1-5.209-5.208"/></g>`), g.Group(children),
	)
}

func ReceiptClese(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 14.583L22.917 25m0-10.417L31.25 25"/><path stroke="#306CFE" d="M43.75 8.333v30.209a5.208 5.208 0 1 1-10.417 0v-5.209H10.417v-25A2.083 2.083 0 0 1 12.5 6.25h29.167a2.083 2.083 0 0 1 2.083 2.083M33.333 38.542v-5.209H6.25v5.209a5.21 5.21 0 0 0 5.208 5.208h27.084a5.21 5.21 0 0 1-5.209-5.208"/></g>`), g.Group(children),
	)
}

func ReceiptDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 20.833L27.083 25l4.167-4.167M27.083 25V14.583"/><path stroke="#306CFE" d="M43.75 8.333v30.209a5.208 5.208 0 1 1-10.417 0v-5.209H10.417v-25A2.083 2.083 0 0 1 12.5 6.25h29.167a2.083 2.083 0 0 1 2.083 2.083M33.333 38.542v-5.209H6.25v5.209a5.21 5.21 0 0 0 5.208 5.208h27.084a5.21 5.21 0 0 1-5.209-5.208"/></g>`), g.Group(children),
	)
}

func ReceiptRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M32.292 19.792H21.875"/><path stroke="#306CFE" d="M43.75 8.333v30.209a5.208 5.208 0 1 1-10.417 0v-5.209H10.417v-25A2.083 2.083 0 0 1 12.5 6.25h29.167a2.083 2.083 0 0 1 2.083 2.083M33.333 38.542v-5.209H6.25v5.209a5.21 5.21 0 0 0 5.208 5.208h27.084a5.21 5.21 0 0 1-5.209-5.208"/></g>`), g.Group(children),
	)
}

func Recycle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m11.167 28.438l-4.605 7.77a2.25 2.25 0 0 0 1.959 3.375h14.396m10.416 0l8.188.105a2.23 2.23 0 0 0 1.916-3.376l-7.041-12.229m-4.958-8.813l-4.355-7.895a2.25 2.25 0 0 0-3.895 0l-7.209 12.5"/><path stroke="#344054" d="m19.375 35.417l3.542 4.166l-4.167 4.167M34.417 29l1.833-5.146l5.688 1.521m-20.646-6.5l-5.375.98l-1.521-5.688"/></g>`), g.Group(children),
	)
}

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 12.5H8.333a2.083 2.083 0 0 0-2.083 2.083v20.834A2.083 2.083 0 0 0 8.333 37.5h20.834m8.333 0h4.167a2.083 2.083 0 0 0 2.083-2.083V14.583a2.083 2.083 0 0 0-2.083-2.083H20.833"/><path stroke="#344054" d="m25 33.333l4.167 4.167L25 41.667m0-25L20.833 12.5L25 8.333"/></g>`), g.Group(children),
	)
}

func RefreshRound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M29.167 37.5H18.75a12.5 12.5 0 0 1-7.208-22.687m9.291-2.313H31.25a12.5 12.5 0 0 1 7.208 22.688"/><path stroke="#344054" d="m25 33.333l4.167 4.167L25 41.667m0-25L20.833 12.5L25 8.333"/></g>`), g.Group(children),
	)
}

func RejectedFileTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m21.875 21.875l6.25 6.25m0-6.25l-6.25 6.25"/><path stroke="#306CFE" d="M14.583 6.25H37.5a2.083 2.083 0 0 1 2.083 2.083v33.334A2.083 2.083 0 0 1 37.5 43.75h-25a2.083 2.083 0 0 1-2.083-2.083v-31.25zm-4.166 4.167h4.166V6.25z"/></g>`), g.Group(children),
	)
}

func Research(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 25v10.417zm-8.334 10.417V31.25zM6.25 31.25l5.896-5.896zm10.417-16.667a6.25 6.25 0 1 0 0 12.5a6.25 6.25 0 0 0 0-12.5"/><path stroke="#306CFE" d="M16.667 6.25h25a2.083 2.083 0 0 1 2.083 2.083v33.334a2.083 2.083 0 0 1-2.083 2.083h-25a2.083 2.083 0 0 1-2.084-2.083v-6.25"/></g>`), g.Group(children),
	)
}

func ResearchPresentationLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 16.667H16.667M31.604 37.75l5.896 6zm-4.52-10.667a6.25 6.25 0 1 1-.001 12.501a6.25 6.25 0 0 1 0-12.5"/><path stroke="#306CFE" d="M41.667 31.25v-25H8.333v25a2.083 2.083 0 0 0 2.084 2.083H12.5M43.75 6.25H6.25"/></g>`), g.Group(children),
	)
}

func Reward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m29.27 35.625l4.063 8.125l4.125-4.23l6.292.063l-5.854-11.625v-.02M20.73 35.625l-4.063 8.125l-4.125-4.23l-6.292.063l5.854-11.625v-.02"/><path stroke="#306CFE" d="M40.167 20.833c0 1.584-1.396 2.917-1.855 4.334c-.458 1.416-.145 3.354-1.041 4.583s-2.813 1.5-3.938 2.417s-2.083 2.625-3.541 3.104c-1.459.479-3.105-.438-4.688-.438s-3.27.896-4.687.438c-1.417-.459-2.313-2.23-3.542-3.104c-1.23-.875-3.146-1.188-4.042-2.417c-.896-1.23-.562-3.125-1.041-4.583c-.48-1.459-1.959-2.75-1.959-4.334s1.396-2.916 1.854-4.333s.146-3.354 1.042-4.583s2.813-1.5 4.042-2.417s1.979-2.625 3.541-3.104c1.563-.48 3.105.437 4.688.437s3.27-.896 4.688-.437C31.104 6.854 32 8.625 33.333 9.5s3.146 1.187 4.042 2.417s.563 3.125 1.042 4.583s1.75 2.75 1.75 4.333M25 14.583a6.25 6.25 0 1 0 0 12.501a6.25 6.25 0 0 0 0-12.5"/></g>`), g.Group(children),
	)
}

func RightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m22.917 31.25l6.25-6.25l-6.25-6.25"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func RightCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 25h16.666m-5.208-5.208L33.333 25l-5.208 5.208"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func RightDirection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.875 20.833H8.333a2.083 2.083 0 0 0-2.083 2.084v4.166a2.083 2.083 0 0 0 2.083 2.084h21.542"/><path stroke="#306CFE" d="m29.875 29.167l-6.083 7.083l6.25 5.417l13.125-15.313a2.084 2.084 0 0 0 0-2.708L30.125 8.333l-6.25 5.417l6 7.083"/></g>`), g.Group(children),
	)
}

func RightDirectionCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 29.167H14.583v-8.334H25v-4.166l9.604 6.625a2.083 2.083 0 0 1 0 3.416L25 33.333z"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func RightDirectionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 29.167H14.583v-8.334H25v-4.166l9.604 6.625a2.083 2.083 0 0 1 0 3.416L25 33.333z"/><path stroke="#306CFE" d="M8.333 43.75h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083"/></g>`), g.Group(children),
	)
}

func RightLeftScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 12.5v25"/><path stroke="#306CFE" d="M12.5 29.167L16.667 25L12.5 20.833m25 0L33.333 25l4.167 4.167M43.75 25H33.333M6.25 25h10.417z"/></g>`), g.Group(children),
	)
}

func RightSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 6.25L20.833 25L10.417 43.75"/><path stroke="#306CFE" d="M10.417 43.75h18.75L39.583 25L29.167 6.25h-18.75"/></g>`), g.Group(children),
	)
}

func RightSquareThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 25H18.75m8.333 4.167L31.25 25l-4.167-4.167"/><path stroke="#306CFE" d="m23.513 6.866l-16.66 16.66a2.083 2.083 0 0 0 0 2.947l16.66 16.661a2.083 2.083 0 0 0 2.947 0l16.66-16.66a2.083 2.083 0 0 0 0-2.947L26.46 6.866a2.083 2.083 0 0 0-2.947 0"/></g>`), g.Group(children),
	)
}

func RotateLock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 29.167h12.5a2.083 2.083 0 0 0 2.083-2.084V18.75a2.083 2.083 0 0 0-2.083-2.083h-12.5a2.083 2.083 0 0 0-2.084 2.083v8.333a2.083 2.083 0 0 0 2.084 2.084M31.25 12.5v4.167h8.333V12.5a4.167 4.167 0 0 0-4.166-4.167v0A4.167 4.167 0 0 0 31.25 12.5"/><path stroke="#306CFE" d="M43.75 37.5v2.083a2.084 2.084 0 0 1-2.083 2.084H8.333a2.083 2.083 0 0 1-2.083-2.084V18.75a2.083 2.083 0 0 1 2.083-2.083H18.75"/><path stroke="#306CFE" d="M6.25 25v8.333l4.167-1.041v-6.25z"/></g>`), g.Group(children),
	)
}

func RouterFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="M33.333 18.75V8.333m6.25 35.417H10.417a2.083 2.083 0 0 1-2.084-2.083V20.833a2.083 2.083 0 0 1 2.084-2.083h29.166a2.083 2.083 0 0 1 2.084 2.083v20.834a2.084 2.084 0 0 1-2.084 2.083"/><path stroke="#344054" stroke-width="2" d="M16.667 27.083h4.166"/><path stroke="#344054" stroke-width="3" d="M33.23 7.292h.208"/></g>`), g.Group(children),
	)
}

func RulerSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 39.583v4.167M6.25 14.583h4.167zM6.25 25h4.167zm29.167 14.583v4.167z"/><path stroke="#306CFE" d="M18.75 31.25L6.854 43.146M43.75 31.25v12.5H8.333a2.083 2.083 0 0 1-2.083-2.083V6.25h12.5v25z"/></g>`), g.Group(children),
	)
}

func RulerSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m22.063 13.208l-2.959 2.959m17.688 11.77l-2.959 2.959zm-7.375-7.354l-2.938 2.938z"/><path stroke="#306CFE" d="m15.087 6.247l-8.839 8.839l28.667 28.667l8.84-8.839z"/></g>`), g.Group(children),
	)
}

func RulerTen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 10.417V6.25m18.75 29.167h-4.167zm0-10.417h-4.167zM14.583 10.417V6.25z"/><path stroke="#306CFE" d="M31.25 18.75L43.146 6.854M6.25 18.75V6.25h35.417a2.083 2.083 0 0 1 2.083 2.083V43.75h-12.5v-25z"/></g>`), g.Group(children),
	)
}

func RulerTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M27.083 35.417h4.167m-4.167-20.834h4.167zm0 10.417h4.167z"/><path stroke="#306CFE" d="M18.75 43.75h12.5V6.25h-12.5z"/></g>`), g.Group(children),
	)
}

func Safebox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 22.917h-2.083m14.583 0a6.25 6.25 0 1 0-12.5 0a6.25 6.25 0 0 0 12.5 0M25 16.667v-2.084zm6.25 6.25h2.083zM25 29.167v2.083z"/><path stroke="#306CFE" d="M41.667 39.583H8.333A2.083 2.083 0 0 1 6.25 37.5V8.333A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083V37.5a2.083 2.083 0 0 1-2.083 2.083m-25 0H12.5v4.167h4.167zm16.666 4.167H37.5v-4.167h-4.167z"/></g>`), g.Group(children),
	)
}

func SafeboxTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 22.917h4.166m-4.166 0a4.167 4.167 0 1 0-8.334 0a4.167 4.167 0 0 0 8.334 0"/><path stroke="#306CFE" d="M41.667 39.583H8.333A2.083 2.083 0 0 1 6.25 37.5V8.333A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083V37.5a2.083 2.083 0 0 1-2.083 2.083m-25 0H12.5v4.167h4.167zm16.666 4.167H37.5v-4.167h-4.167z"/></g>`), g.Group(children),
	)
}

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M12.583 32.167L22.917 12.5m16.458 16.667A25.69 25.69 0 0 0 22.917 8.333v22.5"/><path stroke="#306CFE" d="M28.48 41.667H9.5a2.08 2.08 0 0 1-2.083-1.605L6.25 35.418a2.083 2.083 0 0 1 1.813-2.688l33.333-3.604a2.082 2.082 0 0 1 2.083 3.146a16.94 16.94 0 0 1-15 9.396"/></g>`), g.Group(children),
	)
}

func SatelliteDish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="m31.917 16l7.666-7.667"/><path stroke="#344054" stroke-width="3" d="M39.688 8.333h-.209"/><path stroke="#306CFE" stroke-width="2" d="m12.5 43.75l7.083-18.23m8.709 5.293l5.041 12.937m-25 0H37.5M22.917 6.896a2.21 2.21 0 0 0-3.313.229a15.27 15.27 0 0 0 21.188 21.292a2.21 2.21 0 0 0 .229-3.313z"/></g>`), g.Group(children),
	)
}

func SaveErrorLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M41.667 27.083v6.25m-12.5-18.75h4.166z"/><path stroke="#344054" stroke-width="3" d="M41.563 42.708h.208"/><path stroke="#306CFE" stroke-width="2" d="M33.333 43.75H10.417a2.083 2.083 0 0 1-2.084-2.083V14.583l8.334-8.333h22.916a2.084 2.084 0 0 1 2.084 2.083V18.75"/><path stroke="#306CFE" stroke-width="2" d="M33.333 31.25H16.667v12.5h16.666z"/></g>`), g.Group(children),
	)
}

func ScanAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 29.167v-8.334M25 33.333V16.667zm-8.333-4.166v-8.334z"/><path stroke="#306CFE" d="M16.667 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083v-8.334M33.333 43.75h8.334a2.083 2.083 0 0 0 2.083-2.083v-8.334"/><path stroke="#306CFE" d="M6.25 33.333v8.334a2.083 2.083 0 0 0 2.083 2.083h8.334M6.25 16.667V8.333A2.083 2.083 0 0 1 8.333 6.25h8.334M43.75 16.667V8.333a2.083 2.083 0 0 0-2.083-2.083h-8.334"/></g>`), g.Group(children),
	)
}

func SchoolBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 35.417h8.334zm8.334-27.084a2.083 2.083 0 0 0-2.084-2.083h-4.166a2.083 2.083 0 0 0-2.084 2.083v2.084h8.334zm8.333 31.25h2.083a2.084 2.084 0 0 0 2.084-2.083v-6.25H37.5zm-25-8.333H8.333v6.25a2.083 2.083 0 0 0 2.084 2.083H12.5z"/><path stroke="#306CFE" d="M39.583 18.75V12.5a2.083 2.083 0 0 0-2.083-2.083h-25a2.083 2.083 0 0 0-2.083 2.083v6.25a4.167 4.167 0 0 0 4.166 4.167h20.834a4.167 4.167 0 0 0 4.166-4.167m-4.166 4.167H14.583a4.17 4.17 0 0 1-2.083-.584v19.334a2.083 2.083 0 0 0 2.083 2.083h20.834a2.083 2.083 0 0 0 2.083-2.083V22.333a4.17 4.17 0 0 1-2.083.584"/><path stroke="#344054" d="M25 20.833V25"/></g>`), g.Group(children),
	)
}

func ScreenCapture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/><path stroke="#306CFE" d="M43.75 18.75V8.333a2.083 2.083 0 0 0-2.083-2.083H31.25m0 37.5h10.417a2.083 2.083 0 0 0 2.083-2.083V31.25m-25-25H8.333A2.083 2.083 0 0 0 6.25 8.333V18.75m0 12.5v10.417a2.083 2.083 0 0 0 2.083 2.083H18.75"/></g>`), g.Group(children),
	)
}

func Screenchot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func SearchAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M41.667 41.667L31.146 31.146"/><path stroke="#344054" stroke-width="3" d="m42.708 42.708l-7.291-7.291"/><path stroke="#306CFE" stroke-width="2" d="M20.833 35.417c8.055 0 14.584-6.53 14.584-14.584S28.887 6.25 20.833 6.25S6.25 12.78 6.25 20.833c0 8.055 6.53 14.584 14.583 14.584"/></g>`), g.Group(children),
	)
}

func SearchAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m43.75 43.75l-9.042-9.042"/><path stroke="#306CFE" d="M22.917 39.583c9.204 0 16.666-7.462 16.666-16.666S32.121 6.25 22.917 6.25S6.25 13.712 6.25 22.917c0 9.204 7.462 16.666 16.667 16.666"/></g>`), g.Group(children),
	)
}

func Secure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.75 22.917l4.167 4.166l8.333-8.333"/><path stroke="#306CFE" d="m25 43.75l1.833-.792a22.92 22.92 0 0 0 13.813-19.291l.896-11.5a2.08 2.08 0 0 0-1.584-2.084L25 6.25L10.042 10a2.08 2.08 0 0 0-1.584 2.083l.896 11.5a22.92 22.92 0 0 0 13.813 19.292z"/></g>`), g.Group(children),
	)
}

func SecuredFileFolderTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m37.5 29.167l6.25 2.083c0 8.333-1.292 10.625-6.25 12.5c-4.958-1.875-6.25-4.167-6.25-12.5z"/><path stroke="#306CFE" d="M37.5 20.833V18.75a2.083 2.083 0 0 0-2.083-2.083h-6.25L25 20.833H8.333a2.083 2.083 0 0 0-2.083 2.084V37.5a2.083 2.083 0 0 0 2.083 2.083h14.584"/><path stroke="#306CFE" d="M29.167 16.667L25 20.833H10.417v-12.5A2.083 2.083 0 0 1 12.5 6.25h18.75a2.083 2.083 0 0 1 2.083 2.083v8.334z"/></g>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/><path stroke="#306CFE" d="M41.667 20.833h-1.23a2.08 2.08 0 0 1-1.958-1.395a2.08 2.08 0 0 1 .417-2.375l.854-.855a2.08 2.08 0 0 0 0-2.958l-2.958-2.98a2.084 2.084 0 0 0-2.959 0l-.854.855a2.08 2.08 0 0 1-2.375.417a2.08 2.08 0 0 1-1.437-1.98V8.334a2.083 2.083 0 0 0-2.084-2.083h-4.166a2.083 2.083 0 0 0-2.084 2.083v1.23a2.08 2.08 0 0 1-1.395 1.958v0a2.08 2.08 0 0 1-2.375-.417l-.855-.854a2.08 2.08 0 0 0-2.958 0l-2.98 2.958a2.084 2.084 0 0 0 0 2.959l.855.854a2.08 2.08 0 0 1 .417 2.375a2.08 2.08 0 0 1-1.959 1.396h-1.25a2.083 2.083 0 0 0-2.083 2.083v4.167a2.083 2.083 0 0 0 2.083 2.083h1.23a2.08 2.08 0 0 1 1.958 1.396v0a2.08 2.08 0 0 1-.417 2.375l-.854.854a2.083 2.083 0 0 0 0 2.958l2.938 2.938a2.084 2.084 0 0 0 2.958 0l.854-.854a2.08 2.08 0 0 1 2.375-.417a2.08 2.08 0 0 1 1.396 1.958v1.334a2.083 2.083 0 0 0 2.083 2.083h4.167a2.083 2.083 0 0 0 2.083-2.083v-1.23A2.08 2.08 0 0 1 30.5 38.48a2.08 2.08 0 0 1 2.375.417l.854.854a2.08 2.08 0 0 0 2.959 0l2.937-2.937a2.08 2.08 0 0 0 0-2.959L38.771 33a2.08 2.08 0 0 1-.417-2.375v0a2.08 2.08 0 0 1 1.959-1.396h1.354a2.083 2.083 0 0 0 2.083-2.083v-4.23a2.083 2.083 0 0 0-2.083-2.083"/></g>`), g.Group(children),
	)
}

func SettingsAltFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M33.333 21.542v-4.875h-4.875L25 13.208l-3.458 3.459h-4.875v4.875L13.208 25l3.459 3.458v4.875h4.875L25 36.792l3.458-3.459h4.875v-4.875L36.792 25z"/><path stroke="#344054" stroke-width="3" d="M25.104 25h-.208"/><path stroke="#306CFE" stroke-width="2" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func SettingsAltFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M42.708 42.708a4.167 4.167 0 0 1-5.895 0L25.958 31.854A13.73 13.73 0 0 1 16.667 33A13.54 13.54 0 0 1 6.25 21a13.4 13.4 0 0 1 1.354-7.23l9.063 9.147l5.208-1.042l1.042-5.208l-9.146-9A13.4 13.4 0 0 1 21 6.25a13.54 13.54 0 0 1 12 10.417a13.73 13.73 0 0 1-1.146 9.208l10.854 10.854a4.17 4.17 0 0 1 0 5.98"/>`), g.Group(children),
	)
}

func ShareAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.333 6.25l10.417 8.333l-10.417 8.334V18.75s-10.416 0-14.583 6.25c0 0 2.083-12.5 14.583-14.583z"/><path stroke="#306CFE" d="M43.75 27.083v12.5a2.083 2.083 0 0 1-2.083 2.084H8.333a2.083 2.083 0 0 1-2.083-2.084V12.5a2.083 2.083 0 0 1 2.083-2.083h8.334"/></g>`), g.Group(children),
	)
}

func ShareAltFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m18.083 22.188l13.792-6.875M18.083 27.792l13.792 6.895z"/><path stroke="#306CFE" d="M12.5 18.75a6.25 6.25 0 1 1 0 12.5a6.25 6.25 0 0 1 0-12.5M31.25 37.5a6.25 6.25 0 1 0 12.5 0a6.25 6.25 0 0 0-12.5 0m6.25-18.75a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5"/></g>`), g.Group(children),
	)
}

func ShippingBoxTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 35.417V37.5a2.083 2.083 0 0 1-2.084 2.083h-25A2.083 2.083 0 0 1 6.25 37.5v-25a2.083 2.083 0 0 1 2.083-2.083h25"/><path stroke="#344054" d="M33.333 27.083H43.75M16.667 10.417v12.5l4.166-2.084L25 22.917v-12.5zM43.75 18.75h-8.333z"/></g>`), g.Group(children),
	)
}

func SignInDoubleArrowTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 31.25L14.583 25l6.25-6.25m10.417 12.5L25 25l6.25-6.25M25 25h18.75"/><path stroke="#306CFE" d="M14.583 8.333h-6.25a2.083 2.083 0 0 0-2.083 2.084v29.166a2.083 2.083 0 0 0 2.083 2.084h6.25"/></g>`), g.Group(children),
	)
}

func SignOutAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 25h-25m8.333-8.333L18.75 25l8.333 8.333"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func SignOutLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 8.333h6.25a2.083 2.083 0 0 1 2.083 2.084v29.166a2.083 2.083 0 0 1-2.083 2.084h-6.25"/><path stroke="#306CFE" d="m19.667 11.77l3.25 2.813a2.084 2.084 0 0 1 .229 2.938l-3.021 3.312h15.292v8.334H20.125l2.938 3.416a2.083 2.083 0 0 1-.146 2.938l-3.167 2.708A2.083 2.083 0 0 1 16.667 38L6.75 26.354a2.08 2.08 0 0 1 0-2.708L16.667 12a2.084 2.084 0 0 1 3-.23"/></g>`), g.Group(children),
	)
}

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#306CFE" stroke-width="2" d="M25 18.75v25"/><path stroke="#344054" stroke-width="2" d="M32.375 12.146a10.417 10.417 0 0 1 0 14.729m-14.75 0a10.417 10.417 0 0 1 0-14.73M38.25 6.25a18.75 18.75 0 0 1 0 26.52m-26.5 0a18.75 18.75 0 0 1 0-26.52"/><path stroke="#344054" stroke-width="3" d="M24.896 19.5h.208"/></g>`), g.Group(children),
	)
}

func SignalTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M24.896 25h.208"/><path stroke="#306CFE" stroke-width="2" d="M11.75 38.25a18.75 18.75 0 0 1 0-26.5m5.875 5.875a10.417 10.417 0 0 0 0 14.75M38.25 38.25a18.75 18.75 0 0 0 0-26.5m-5.875 20.625a10.415 10.415 0 0 0 0-14.75"/></g>`), g.Group(children),
	)
}

func SleepingBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 35.417h4.166m6.25 8.333H16.667a2.083 2.083 0 0 1-2.084-2.083V14.583a8.333 8.333 0 0 1 8.334-8.333h4.166a8.333 8.333 0 0 1 8.334 8.333v27.084a2.083 2.083 0 0 1-2.084 2.083M22.917 27.083h4.166z"/><path stroke="#306CFE" d="M35.417 18.75v22.917a2.083 2.083 0 0 1-2.084 2.083H16.667a2.083 2.083 0 0 1-2.084-2.083V18.75z"/></g>`), g.Group(children),
	)
}

func SleepingBagTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 6.25h4.166a8.333 8.333 0 0 1 8.334 8.333v27.084a2.083 2.083 0 0 1-2.084 2.083H16.667a2.083 2.083 0 0 1-2.084-2.083V14.583a8.333 8.333 0 0 1 8.334-8.333"/><path stroke="#306CFE" d="M25 43.75V22.917a4.167 4.167 0 0 1 4.167-4.167h6.25v22.917a2.083 2.083 0 0 1-2.084 2.083zm-10.417-2.083V18.75h6.25A4.167 4.167 0 0 1 25 22.917V43.75h-8.333a2.083 2.083 0 0 1-2.084-2.083"/></g>`), g.Group(children),
	)
}

func Sofa(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 18.75V12.5a2.083 2.083 0 0 1 2.083-2.083h25a2.083 2.083 0 0 1 2.083 2.083v6.25M10.417 39.583v-4.166m29.166 4.166v-4.166z"/><path stroke="#306CFE" d="M43.75 22.917v8.333a4.167 4.167 0 0 1-4.167 4.167H10.417A4.167 4.167 0 0 1 6.25 31.25v-8.333a4.166 4.166 0 0 1 4.875-4.167a4.31 4.31 0 0 1 3.458 4.396v3.937h20.834v-3.937a4.31 4.31 0 0 1 3.458-4.396a4.168 4.168 0 0 1 4.875 4.167"/></g>`), g.Group(children),
	)
}

func SofaTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M10.417 18.75a8.333 8.333 0 0 1 8.333-8.333h12.5a8.333 8.333 0 0 1 8.333 8.333M10.417 39.583v-4.166m29.166 4.166v-4.166z"/><path stroke="#306CFE" d="M43.75 22.917v8.333a4.167 4.167 0 0 1-4.167 4.167H10.417A4.167 4.167 0 0 1 6.25 31.25v-8.333a4.166 4.166 0 0 1 4.875-4.167a4.31 4.31 0 0 1 3.458 4.396v3.937h20.834v-3.937a4.31 4.31 0 0 1 3.458-4.396a4.168 4.168 0 0 1 4.875 4.167"/></g>`), g.Group(children),
	)
}

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 25h29.167"/><path stroke="#306CFE" d="M6.25 35.417h20.833M6.25 14.583h37.5z"/></g>`), g.Group(children),
	)
}

func SortAscending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 31.25L12.5 39.583L4.167 31.25m8.333 8.333V8.333"/><path stroke="#306CFE" d="M41.667 12.5H20.833m20.834 20.833H31.25zm0-10.416H27.083z"/></g>`), g.Group(children),
	)
}

func SortDescending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M4.167 16.667L12.5 8.333l8.333 8.334M12.5 8.333v31.25"/><path stroke="#306CFE" d="M41.667 35.417H20.833m20.834-20.834H31.25zm0 10.417H27.083z"/></g>`), g.Group(children),
	)
}

func SoundIncrease(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 25h12.5m-6.25 6.25v-12.5z"/><path stroke="#306CFE" d="M22.917 10.417v29.166l-8.334-8.333h-6.25a2.083 2.083 0 0 1-2.083-2.083v-8.334a2.083 2.083 0 0 1 2.083-2.083h6.25z"/></g>`), g.Group(children),
	)
}

func SoundMax(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M38.25 11.75a18.75 18.75 0 0 1 0 26.5m-5.875-20.625a10.415 10.415 0 0 1 0 14.75"/><path stroke="#306CFE" d="M22.917 10.417v29.166l-8.334-8.333h-6.25a2.083 2.083 0 0 1-2.083-2.083v-8.334a2.083 2.083 0 0 1 2.083-2.083h6.25z"/></g>`), g.Group(children),
	)
}

func SoundMin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M36.542 17.625a10.415 10.415 0 0 1 0 14.75"/><path stroke="#306CFE" d="M27.083 10.417v29.166L18.75 31.25H12.5a2.083 2.083 0 0 1-2.083-2.083v-8.334A2.083 2.083 0 0 1 12.5 18.75h6.25z"/></g>`), g.Group(children),
	)
}

func SoundReduce(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 25h12.5"/><path stroke="#306CFE" d="M22.917 10.417v29.166l-8.334-8.333h-6.25a2.083 2.083 0 0 1-2.083-2.083v-8.334a2.083 2.083 0 0 1 2.083-2.083h6.25z"/></g>`), g.Group(children),
	)
}

func SquareRoot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.25 25h2.73a2.08 2.08 0 0 1 1.937 1.354l5.75 15.313l5.937-31.625a2.08 2.08 0 0 1 2.084-1.709H43.75"/>`), g.Group(children),
	)
}

func SquareRootOfX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m29.167 25l8.333 8.333m0-8.333l-8.333 8.333"/><path stroke="#306CFE" d="M6.25 25h2.98a2.08 2.08 0 0 1 1.79 1.02l5.647 9.397l5.812-19.354a2.08 2.08 0 0 1 2.084-1.48H43.75"/></g>`), g.Group(children),
	)
}

func SquareRootSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M14.583 25h2.084l4.166 8.333L25 16.667h10.417"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 37.5h16.666m-16.666-25h16.666zm0 8.333h16.666zm0 8.334h16.666z"/><path stroke="#306CFE" d="M33.333 6.25v37.5M16.667 6.25v37.5z"/></g>`), g.Group(children),
	)
}

func StairsThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M41.667 18.75L8.333 6.25m31.25 11.708v17.459zM10.417 7.042v15.875z"/><path stroke="#306CFE" d="M6.25 43.75h37.5M8.333 22.917V43.75h33.334v-8.333h-8.334V31.25H25v-4.167h-8.333v-4.166z"/></g>`), g.Group(children),
	)
}

func StampThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 41.667H16.667v2.083h16.666z"/><path stroke="#306CFE" d="M20.833 29.167h8.334a26.1 26.1 0 0 1 3.208-10.792A8.333 8.333 0 1 0 16.917 12.5a8.33 8.33 0 0 0 .687 5.917a26.9 26.9 0 0 1 3.23 10.75m13.271 1.729a2.084 2.084 0 0 0-2.083-1.73H18.063a2.08 2.08 0 0 0-2.084 1.73l-1.396 8.333a2.084 2.084 0 0 0 2.084 2.438h16.666a2.083 2.083 0 0 0 2.084-2.438z"/></g>`), g.Group(children),
	)
}

func Stats(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 35.417V14.583M25 35.417v-12.5zm-8.333 0V31.25z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func StickyNotesNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 14.583v27.084a2.083 2.083 0 0 1-2.084 2.083H18.75l-12.5-12.5V14.583A2.083 2.083 0 0 1 8.333 12.5h25"/><path stroke="#306CFE" d="m6.25 31.25l12.5 12.5v-12.5z"/><path stroke="#344054" d="m34.563 13.354l-5.396 5.396m12.5-8.333a4.167 4.167 0 1 0-8.334 0v0a4.167 4.167 0 1 0 8.334 0"/></g>`), g.Group(children),
	)
}

func StickyNotesSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 14.583v18.75a2.083 2.083 0 0 1-2.084 2.084h-25a2.083 2.083 0 0 1-2.083-2.084v-18.75A2.083 2.083 0 0 1 8.333 12.5h25"/><path stroke="#306CFE" d="M10.417 43.75h31.25a2.083 2.083 0 0 0 2.083-2.083V20.833"/><path stroke="#344054" d="m34.563 13.354l-5.396 5.396m12.5-8.333a4.167 4.167 0 1 0-8.334 0v0a4.167 4.167 0 1 0 8.334 0"/></g>`), g.Group(children),
	)
}

func StickyNotesTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M29.167 10.417h10.416a2.083 2.083 0 0 1 2.084 2.083v18.75l-12.5 12.5h-18.75a2.083 2.083 0 0 1-2.084-2.083V12.5a2.083 2.083 0 0 1 2.084-2.083h10.416"/><path stroke="#306CFE" d="M29.167 31.25v12.5l12.5-12.5z"/><path stroke="#344054" d="M29.167 16.667V8.333c0-1.15-.933-2.083-2.084-2.083h-4.166a2.084 2.084 0 0 0-2.084 2.083v8.334c0 1.15.933 2.083 2.084 2.083h4.166c1.15 0 2.084-.933 2.084-2.083"/></g>`), g.Group(children),
	)
}

func Stool(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M13.542 29.167h22.916"/><path stroke="#306CFE" d="M35.417 14.583L37.5 43.75M14.583 14.583L12.5 43.75m2.083-37.5h20.834A2.083 2.083 0 0 1 37.5 8.333v6.25h-25v-6.25a2.083 2.083 0 0 1 2.083-2.083"/></g>`), g.Group(children),
	)
}

func StoolTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 27.083h6.25A2.083 2.083 0 0 0 33.333 25v-2.083"/><path stroke="#306CFE" d="M18.75 43.75h12.5m6.25-29.167h-25v-6.25a2.083 2.083 0 0 1 2.083-2.083h20.834A2.083 2.083 0 0 1 37.5 8.333zm-12.5 0V43.75z"/></g>`), g.Group(children),
	)
}

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/>`), g.Group(children),
	)
}

func StopwatchSevenSecond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 6.25h-12.5m-4.792 10.833l-3.125-3.125zm22.084 0l3.125-3.125zM17.708 28.125a7.25 7.25 0 0 0 3 5.875L25 29.167v-8.334a7.29 7.29 0 0 0-7.292 7.292M25 12.5V6.25z"/><path stroke="#306CFE" d="M25 43.75c8.63 0 15.625-6.996 15.625-15.625c0-8.63-6.996-15.625-15.625-15.625c-8.63 0-15.625 6.996-15.625 15.625c0 8.63 6.996 15.625 15.625 15.625"/></g>`), g.Group(children),
	)
}

func StopwatchThreeSecond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 6.25h12.5m4.792 10.833l3.125-3.125zm-22.084 0l-3.125-3.125zM25 20.833v8.334h7.208a5.7 5.7 0 0 0 0-1.042A7.29 7.29 0 0 0 25 20.833m0-8.333V6.25z"/><path stroke="#306CFE" d="M25 43.75c8.63 0 15.625-6.996 15.625-15.625c0-8.63-6.996-15.625-15.625-15.625c-8.63 0-15.625 6.996-15.625 15.625c0 8.63 6.996 15.625 15.625 15.625"/></g>`), g.Group(children),
	)
}

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 39.583h16.666M18.75 31.25l-4.167 12.5zm16.667 12.5l-4.167-12.5zM20.833 22.917h8.334"/><path stroke="#306CFE" d="M20.833 6.25h8.334a8.333 8.333 0 0 1 8.333 8.333v14.584a2.083 2.083 0 0 1-2.083 2.083H14.583a2.083 2.083 0 0 1-2.083-2.083V14.583a8.333 8.333 0 0 1 8.333-8.333"/></g>`), g.Group(children),
	)
}

func Suit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 12.5v29.167a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V12.5a2.083 2.083 0 0 1 2.083-2.083H12.5l4.167-4.167h16.666l4.167 4.167h4.167A2.083 2.083 0 0 1 43.75 12.5"/><path stroke="#344054" d="M16.667 6.25L25 15.188l8.333-8.938zm16.666 0L25 15.188l8.333 5.645L37.5 10.417zM12.5 10.417l4.167 10.416L25 15.188L16.667 6.25z"/></g>`), g.Group(children),
	)
}

func SuitcaseBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 6.25h8.334a6.25 6.25 0 0 1 6.25 6.25v2.083H14.583V12.5a6.25 6.25 0 0 1 6.25-6.25M22.917 25h4.166"/><path stroke="#306CFE" d="m41.313 29.98l-1.5 11.937a2.084 2.084 0 0 1-2.084 1.833H12.25a2.084 2.084 0 0 1-2.083-1.833l-1.5-11.938"/><path stroke="#306CFE" d="m42.333 29.646l-9.437 3.146a25 25 0 0 1-15.792 0l-9.437-3.146a2.084 2.084 0 0 1-1.417-2.084V14.584h37.5v13.084a2.08 2.08 0 0 1-1.417 1.979"/></g>`), g.Group(children),
	)
}

func Suspension(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 43.75a4.167 4.167 0 1 1 0-8.334a4.167 4.167 0 0 1 0 8.334m0-8.333v-25zM27.083 6.25h-4.166v4.167h4.166z"/><path stroke="#344054" d="m20.833 17.708l8.334-2.083m-8.334 8.333l8.334-2.083zm0 6.25l8.334-2.083z"/></g>`), g.Group(children),
	)
}

func SuspensionTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M20.833 39.583a4.166 4.166 0 1 1 8.332 0a4.166 4.166 0 0 1-8.332 0M25 10.417v25zm-2.083 0h4.166V6.25h-4.166z"/><path stroke="#344054" d="m20.833 15.625l8.334 2.083m-8.334 4.167l8.334 2.083zm0 6.25l8.334 2.083z"/></g>`), g.Group(children),
	)
}

func SwissArmyKnife(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m22.917 23.896l4.875 4.896a8.334 8.334 0 0 0 11.791 0L22.917 12.125q.015.188 0 .375z"/><path stroke="#306CFE" d="M22.917 12.5a6.25 6.25 0 1 0-12.5 0v25a6.25 6.25 0 1 0 12.5 0z"/></g>`), g.Group(children),
	)
}

func SwitchDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M29.167 13.542a7.17 7.17 0 0 0 2.208 5.208H11.458a5.209 5.209 0 0 1 0-10.417h19.917a7.17 7.17 0 0 0-2.208 5.209"/><path stroke="#344054" d="M43.75 12.917a7.27 7.27 0 0 0-12.375-4.584a7.292 7.292 0 0 0 5.083 12.5a7.27 7.27 0 0 0 7.292-6.666q.03-.313 0-.625q.03-.313 0-.625"/><path stroke="#306CFE" d="M20.833 36.458a7.17 7.17 0 0 0-2.208-5.208h19.917a5.208 5.208 0 1 1 0 10.417H18.625a7.17 7.17 0 0 0 2.208-5.209"/><path stroke="#344054" d="M6.25 37.083a7.27 7.27 0 0 0 12.375 4.584a7.291 7.291 0 0 0-5.083-12.5a7.27 7.27 0 0 0-7.292 6.666a4.7 4.7 0 0 0 0 1.25"/></g>`), g.Group(children),
	)
}

func SwitchLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M19.792 30.208a5.208 5.208 0 1 0 0-10.416a5.208 5.208 0 0 0 0 10.416"/><path stroke="#306CFE" d="M31.25 37.5c6.904 0 12.5-5.596 12.5-12.5s-5.596-12.5-12.5-12.5h-12.5c-6.904 0-12.5 5.596-12.5 12.5s5.596 12.5 12.5 12.5z"/></g>`), g.Group(children),
	)
}

func SwitchRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M30.208 30.208a5.208 5.208 0 1 0 0-10.416a5.208 5.208 0 0 0 0 10.416"/><path stroke="#306CFE" d="M31.25 12.5h-12.5c-6.904 0-12.5 5.596-12.5 12.5s5.596 12.5 12.5 12.5h12.5c6.904 0 12.5-5.596 12.5-12.5s-5.596-12.5-12.5-12.5"/></g>`), g.Group(children),
	)
}

func TableLamp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 31.25v12.5m-6.25 0h12.5z"/><path stroke="#344054" d="M37.5 31.25L34.188 8.042a2.084 2.084 0 0 0-2.084-1.792H17.875a2.08 2.08 0 0 0-2.083 1.792L12.5 31.25z"/></g>`), g.Group(children),
	)
}

func TabletLaptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M18.75 31.25H6.25v2.083a2.083 2.083 0 0 0 2.083 2.084H18.75zm0-14.583a2.083 2.083 0 0 1 2.083-2.084h20.834v-6.25a2.083 2.083 0 0 0-2.084-2.083H10.417a2.083 2.083 0 0 0-2.084 2.083V31.25H18.75z"/><path stroke="#344054" d="M41.667 14.583H20.833c-1.15 0-2.083.933-2.083 2.084v25c0 1.15.933 2.083 2.083 2.083h20.834c1.15 0 2.083-.933 2.083-2.083v-25c0-1.15-.933-2.084-2.083-2.084"/></g>`), g.Group(children),
	)
}

func TakeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 43.75h37.5"/><path stroke="#306CFE" d="m42.604 14.333l-9.646 4.604a2.08 2.08 0 0 0-1.062 1.125l-2.5 6.417c-.151.398-.42.74-.771.98l-5.708 3.79l.937-7.978l-7.562 3.604a2.29 2.29 0 0 1-2.605-.48L6.25 18.209l3.77-2.354a2.23 2.23 0 0 1 2.48.063l2.708 1.875L31.5 7.604a8.33 8.33 0 0 1 4.563-1.354a7.67 7.67 0 0 1 7.062 4.5l.48 1.167a1.875 1.875 0 0 1-1 2.416"/></g>`), g.Group(children),
	)
}

func Tax(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M43.75 12.5L31.25 25m-2.083 10.417H14.583zm-6.25-8.334h-8.334z"/><path stroke="#344054" stroke-width="3" d="M32.188 13.542h.208m10.208 10.416h.209"/><path stroke="#306CFE" stroke-width="2" d="M37.5 33.333v8.334a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V8.333A2.083 2.083 0 0 1 8.333 6.25h14.584"/></g>`), g.Group(children),
	)
}

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m27.083 25l6.25 18.75M27.083 25l-6.25 18.75z"/><path stroke="#306CFE" d="M43.75 22.354L39.438 6.25m3.77 14.083l-22.125 5.938a2.08 2.08 0 0 1-2.562-1.48l-2.146-8.124a2.083 2.083 0 0 1 1.458-2.563l22.146-5.77zM8.333 27.5l10.105-2.708l-2.063-8.125L6.25 19.458z"/></g>`), g.Group(children),
	)
}

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m26.938 27.75l6.395 16H16.667l6.395-16a2.082 2.082 0 0 1 3.875 0"/><path stroke="#306CFE" d="M25 6.25a42 42 0 0 1-14.583 11.208L8.52 41.5a2.083 2.083 0 0 0 2.083 2.25h28.813A2.085 2.085 0 0 0 41.5 41.5l-1.917-24.042A42 42 0 0 1 25 6.25"/><path stroke="#344054" d="M6.25 18.75C16.667 16.667 25 6.25 25 6.25s8.333 10.417 18.75 12.5"/></g>`), g.Group(children),
	)
}

func ThreeRdPlace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 16.667h7.146a1.21 1.21 0 0 1 1.188 1.187v14.292a1.21 1.21 0 0 1-1.188 1.187h-7.146M22.917 25h6.25"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 27.083v-4.166m0-12.5v4.166zm0 25v4.166z"/><path stroke="#306CFE" d="M37.5 25a6.25 6.25 0 0 0 6.25 6.25v6.25a2.083 2.083 0 0 1-2.083 2.083H8.333A2.083 2.083 0 0 1 6.25 37.5v-6.25a6.25 6.25 0 0 0 0-12.5V12.5a2.083 2.083 0 0 1 2.083-2.083h33.334A2.083 2.083 0 0 1 43.75 12.5v6.25A6.25 6.25 0 0 0 37.5 25"/></g>`), g.Group(children),
	)
}

func TimeClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 14.583V25h-8.333"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v33.334c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func TimerError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M25.104 34.375h-.208"/><path stroke="#344054" stroke-width="2" d="M18.75 6.25h12.5M25 25v-4.167zm0-12.5V6.25z"/><path stroke="#306CFE" stroke-width="2" d="m36.042 17.083l3.125-3.125zm-22.084 0l-3.125-3.125zM9.375 28.125a15.625 15.625 0 1 0 31.25 0a15.625 15.625 0 0 0-31.25 0"/></g>`), g.Group(children),
	)
}

func TimerFiveSecond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 20.833v8.334l2.917 3.27M18.75 6.25h12.5m4.792 10.833l3.125-3.125zm-22.084 0l-3.125-3.125zM25 12.5V6.25z"/><path stroke="#306CFE" d="M25 43.75c8.63 0 15.625-6.996 15.625-15.625c0-8.63-6.996-15.625-15.625-15.625c-8.63 0-15.625 6.996-15.625 15.625c0 8.63 6.996 15.625 15.625 15.625"/></g>`), g.Group(children),
	)
}

func Toaster(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 35.417h-4.167M12.5 6.25h20.833a4.167 4.167 0 1 1 0 8.333v8.334H12.5v-8.334a4.167 4.167 0 1 1 0-8.333"/><path stroke="#306CFE" d="M39.583 29.167h4.167zm0 12.5V25a2.083 2.083 0 0 0-2.083-2.083H8.333A2.083 2.083 0 0 0 6.25 25v16.667a2.083 2.083 0 0 0 2.083 2.083H37.5a2.083 2.083 0 0 0 2.083-2.083"/></g>`), g.Group(children),
	)
}

func ToasterTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 35.417h4.167m6.25-25a4.167 4.167 0 0 1-4.167 4.166v8.334H16.667v-8.334a4.167 4.167 0 1 1 0-8.333H37.5a4.167 4.167 0 0 1 4.167 4.167"/><path stroke="#306CFE" d="M10.417 29.167H6.25M12.5 43.75h29.167a2.083 2.083 0 0 0 2.083-2.083V25a2.083 2.083 0 0 0-2.083-2.083H12.5A2.083 2.083 0 0 0 10.417 25v16.667A2.083 2.083 0 0 0 12.5 43.75"/></g>`), g.Group(children),
	)
}

func Toilet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m20.833 22.917l12.5-10.417"/><path stroke="#306CFE" d="M33.333 31.25H16.667l-2.084 12.5h20.834zm-12.5-8.333H12.5V8.333a2.083 2.083 0 0 1 2.083-2.083h4.167a2.083 2.083 0 0 1 2.083 2.083zM37.5 27.083v-4.166h-25v4.166a4.167 4.167 0 0 0 4.167 4.167h16.666a4.167 4.167 0 0 0 4.167-4.167"/></g>`), g.Group(children),
	)
}

func ToiletPaper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M14.583 18.854v-.208"/><path stroke="#306CFE" stroke-width="2" d="M14.583 29.167h8.334zm0-20.834C9.98 8.333 6.25 13 6.25 18.75s3.73 10.417 8.333 10.417c4.605 0 8.334-4.667 8.334-10.417s-3.73-10.417-8.334-10.417"/><path stroke="#306CFE" stroke-width="2" d="M22.917 18.75v22.917H43.75V18.75c0-5.75-3.73-10.417-8.333-10.417H14.583"/></g>`), g.Group(children),
	)
}

func ToiletPaperNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M35.417 14.583h-8.334m8.334 12.5V43.75l-3.48-2.083l-3.479 2.083L25 41.667l-3.458 2.083l-3.5-2.083l-3.459 2.083V27.083zm-16.667-12.5h-4.167z"/><path stroke="#306CFE" d="M43.75 25V8.333c0-1.15-.933-2.083-2.083-2.083H8.333c-1.15 0-2.083.933-2.083 2.083V25c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.932 2.083-2.083"/></g>`), g.Group(children),
	)
}

func ToiletPaperSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M18.854 14.583h-.208"/><path stroke="#306CFE" stroke-width="2" d="M29.167 14.583v8.334zM18.75 22.917c5.75 0 10.417-3.73 10.417-8.334S24.5 6.25 18.75 6.25S8.333 9.98 8.333 14.583c0 4.605 4.667 8.334 10.417 8.334"/><path stroke="#306CFE" stroke-width="2" d="M18.75 22.917h22.917l-2.084 3.479l2.084 3.479l-2.084 3.458l2.084 3.459l-2.084 3.479l2.084 3.479H18.75c-5.75 0-10.417-3.73-10.417-8.333V14.583"/></g>`), g.Group(children),
	)
}

func ToiletTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 22.917H14.583V6.25h20.834zm2.083 4.166v-4.166h-25v4.166a4.167 4.167 0 0 0 4.167 4.167h16.666a4.167 4.167 0 0 0 4.167-4.167m-4.167 4.167H16.667l-2.084 12.5h20.834z"/><path stroke="#344054" d="M10.417 22.917h29.166M12.5 6.25h25z"/></g>`), g.Group(children),
	)
}

func TopCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 33.333V16.667m-5.208 5.208L25 16.667l5.208 5.208"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func Touchid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 22.917a33.2 33.2 0 0 1-6.25 20.833m-14.229-7.98c1.73-1.77 3.813-9.29 3.813-12.853A8.333 8.333 0 0 1 25 14.583a8.1 8.1 0 0 1 4.167 1.125"/><path stroke="#306CFE" d="M18.75 7.458A16.9 16.9 0 0 1 25 6.25a16.666 16.666 0 0 1 16.667 16.667c0 2.333 0 10.416-4.167 16.666M6.958 29.167a19.7 19.7 0 0 0 1.375-6.25A16.67 16.67 0 0 1 12.5 11.896M25 22.917c0 4.166-2.083 16.666-8.333 20.833"/></g>`), g.Group(children),
	)
}

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 22.917a16.96 16.96 0 0 0 8.334 6.25"/><path stroke="#344054" d="M20.833 29.167c8.334-2.084 8.334-12.5 8.334-12.5M25 14.583v2.084m-8.333 0h16.666z"/><path stroke="#306CFE" d="M43.75 8.333v25a2.083 2.083 0 0 1-2.083 2.084h-8.334L25 43.75l-8.333-8.333H8.333a2.083 2.083 0 0 1-2.083-2.084v-25A2.083 2.083 0 0 1 8.333 6.25h33.334a2.083 2.083 0 0 1 2.083 2.083"/></g>`), g.Group(children),
	)
}

func TreeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M34.875 25.354c.352.882.535 1.822.542 2.771a7.29 7.29 0 0 1-7.292 7.292a7.2 7.2 0 0 1-3.125-.73a7.2 7.2 0 0 1-3.125.73a7.29 7.29 0 0 1-7.292-7.292c.007-.95.19-1.889.542-2.77a7.25 7.25 0 0 1 1.688-12.21a8.333 8.333 0 0 1 16.375 0a7.25 7.25 0 0 1 1.687 12.21"/><path stroke="#344054" d="M20.833 43.75h8.334m-4.167 0V22.917z"/></g>`), g.Group(children),
	)
}

func TriangleRulerPencil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m30.604 6.854l-2.916 2.917a2.083 2.083 0 0 0 0 2.916l10.229 10.23h5.833v-5.834L33.52 6.854a2.083 2.083 0 0 0-2.916 0"/><path stroke="#306CFE" d="M39.583 43.75H6.25V10.417z"/></g>`), g.Group(children),
	)
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M43.75 12.5v4.167A8.333 8.333 0 0 1 35.854 25c.187-.681.305-1.379.354-2.083l.813-10.417zm-30 9.98l-.77-9.98H6.25v4.167A8.334 8.334 0 0 0 14.23 25a11.3 11.3 0 0 1-.48-2.52"/><path stroke="#306CFE" d="M29.375 32.438a11.2 11.2 0 0 1-5.604.833a10.4 10.4 0 0 1-3.125-.834L18.75 43.75h12.5zM14.75 6.25h20.5a2.084 2.084 0 0 1 2.083 2.25l-1.125 14.417a11.25 11.25 0 0 1-12.5 10.416A11.58 11.58 0 0 1 13.75 22.48L12.667 8.5a2.085 2.085 0 0 1 2.083-2.25"/><path stroke="#344054" d="M35.417 43.75H14.583"/></g>`), g.Group(children),
	)
}

func TrowelThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m23.813 26.188l-6.021 6.041"/><path stroke="#306CFE" d="M7.48 42.52a4.167 4.167 0 0 1 0-5.874l5.874-5.896a2.08 2.08 0 0 1 2.959 0l2.937 2.938a2.083 2.083 0 0 1 0 2.958l-5.896 5.875a4.167 4.167 0 0 1-5.875 0m22.916-9.708a2.083 2.083 0 0 0 2.979 0c3.667-3.833 12.458-14.458 9.958-26.145c-11.687-2.5-22.312 6.291-26.125 10a2.083 2.083 0 0 0 0 2.979z"/></g>`), g.Group(children),
	)
}

func TruckLift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 22.917H31.25m8.333 12.5V6.25zM10.417 6.25v29.167zm0 16.667h8.333z"/><path stroke="#344054" d="M43.75 43.75H6.25V37.5a2.083 2.083 0 0 1 2.083-2.083h33.334A2.083 2.083 0 0 1 43.75 37.5z"/></g>`), g.Group(children),
	)
}

func TurnAroundDownDirection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M12.5 41.667V15.625a9.355 9.355 0 0 1 9.375-9.375v0a9.354 9.354 0 0 1 9.375 9.375V43.75"/><path stroke="#344054" d="m37.5 37.5l-6.25 6.25L25 37.5"/></g>`), g.Group(children),
	)
}

func TurnAroundLeftTopDirectionTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M6.25 18.75h28.125a9.355 9.355 0 0 1 9.375 9.375v0a9.356 9.356 0 0 1-9.375 9.375h-7.292"/><path stroke="#344054" d="m12.5 25l-6.25-6.25l6.25-6.25"/></g>`), g.Group(children),
	)
}

func TurnAroundRightDirectionTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M22.917 37.5h-7.292a9.353 9.353 0 0 1-9.375-9.375a9.354 9.354 0 0 1 9.375-9.375H43.75"/><path stroke="#344054" d="m37.5 12.5l6.25 6.25L37.5 25"/></g>`), g.Group(children),
	)
}

func TurnAroundUpDirection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M37.5 8.333v26.042a9.353 9.353 0 0 1-9.375 9.375v0a9.356 9.356 0 0 1-9.375-9.375V6.25"/><path stroke="#344054" d="m12.5 12.5l6.25-6.25L25 12.5"/></g>`), g.Group(children),
	)
}

func TurnLeftSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M18.75 20.833h14.583a2.083 2.083 0 0 1 2.084 2.084v18.75A2.083 2.083 0 0 0 37.5 43.75h4.167a2.083 2.083 0 0 0 2.083-2.083V20.833a8.333 8.333 0 0 0-8.333-8.333H18.75"/><path stroke="#344054" d="M18.75 12.5V6.25L8.167 15.063a2.083 2.083 0 0 0 0 3.208l10.583 8.812v-6.25"/></g>`), g.Group(children),
	)
}

func TurnRightSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M31.25 20.833H16.667a2.083 2.083 0 0 0-2.084 2.084v18.75A2.083 2.083 0 0 1 12.5 43.75H8.333a2.083 2.083 0 0 1-2.083-2.083V20.833a8.333 8.333 0 0 1 8.333-8.333H31.25"/><path stroke="#344054" d="M31.25 12.5V6.25l10.583 8.813a2.083 2.083 0 0 1 0 3.208L31.25 27.083v-6.25"/></g>`), g.Group(children),
	)
}

func TvStand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m29.167 22.917l2.083 8.333zM18.75 31.25l2.083-8.333zm18.75-8.333h-25a2.083 2.083 0 0 1-2.083-2.084v-12.5A2.083 2.083 0 0 1 12.5 6.25h25a2.083 2.083 0 0 1 2.083 2.083v12.5a2.083 2.083 0 0 1-2.083 2.084"/><path stroke="#344054" d="M39.583 39.583v4.167M6.25 39.583h37.5V31.25H6.25zm0 0H25V31.25H6.25zm4.167 0v4.167z"/></g>`), g.Group(children),
	)
}

func TvStandTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="m18.75 35.417l2.083-8.334m18.75 0H10.417A2.083 2.083 0 0 1 8.333 25V8.333a2.083 2.083 0 0 1 2.084-2.083h29.166a2.083 2.083 0 0 1 2.084 2.083V25a2.083 2.083 0 0 1-2.084 2.083m-10.416 0l2.083 8.334z"/><path stroke="#344054" d="M6.25 43.75h37.5v-8.333H6.25zm0 0H25v-8.333H6.25z"/></g>`), g.Group(children),
	)
}

func TwoK(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m33.708 23.313l10.042 16.27M29.167 10.417v29.166zm0 18.75l14.583-18.75z"/><path stroke="#306CFE" d="M20.833 39.583H6.25a14.35 14.35 0 0 1 7.23-12.5L17.166 25a7.33 7.33 0 0 0 3.666-6.25v-1.042a7.29 7.29 0 0 0-7.291-7.291v0a7.29 7.29 0 0 0-7.292 7.291v1.042"/></g>`), g.Group(children),
	)
}

func TwoNdPlace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M29.167 33.333h-8.334A8.33 8.33 0 0 1 25 26.208L27.083 25a4.17 4.17 0 0 0 2.084-3.604v-.563A4.167 4.167 0 0 0 25 16.667v0a4.167 4.167 0 0 0-4.167 4.166"/><path stroke="#306CFE" d="M25 43.75c10.355 0 18.75-8.395 18.75-18.75S35.355 6.25 25 6.25S6.25 14.645 6.25 25S14.645 43.75 25 43.75"/></g>`), g.Group(children),
	)
}

func UnicycleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M16.667 6.25h1.229c1.883.001 3.69.75 5.02 2.083v0a7.1 7.1 0 0 0 5.022 2.084h3.312"/><path stroke="#344054" d="M25 43.75c5.753 0 10.417-4.664 10.417-10.417S30.753 22.917 25 22.917S14.583 27.58 14.583 33.333S19.247 43.75 25 43.75"/><path stroke="#306CFE" d="M25 33.333V10.417M6.25 39.583a39.3 39.3 0 0 0 17.646 4.167h2.208a39.3 39.3 0 0 0 17.646-4.167"/></g>`), g.Group(children),
	)
}

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 18.75v-4.167A8.333 8.333 0 0 1 25 6.25v0a8.33 8.33 0 0 1 7.208 4.167"/><path stroke="#306CFE" d="M37.5 18.75h-25c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.932 2.083 2.083 2.083h25c1.15 0 2.083-.933 2.083-2.083V20.833c0-1.15-.932-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func UnlockOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M25 35.417v-2.084"/><path stroke="#344054" stroke-width="3" d="M24.896 31.25h.208"/><path stroke="#306CFE" stroke-width="2" d="M35.417 20.833H14.583v-4.166A10.417 10.417 0 0 1 25 6.25v0a10.42 10.42 0 0 1 9.02 5.188m5.563 11.479v5.645A14.98 14.98 0 0 1 25.438 43.75a14.583 14.583 0 0 1-15.021-14.583v-6.25a2.084 2.084 0 0 1 2.083-2.084h25a2.083 2.083 0 0 1 2.083 2.084"/></g>`), g.Group(children),
	)
}

func UpAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M31.25 25v16.667a2.083 2.083 0 0 1-2.083 2.083h-8.334a2.083 2.083 0 0 1-2.083-2.083V25"/><path stroke="#306CFE" d="m18.75 25l-4.604 3.083a2.083 2.083 0 0 1-2.834-.5L7.48 22.48a2.083 2.083 0 0 1 .396-2.896L23.729 7.25a2.08 2.08 0 0 1 2.542 0l15.854 12.333a2.083 2.083 0 0 1 .396 2.896l-3.834 5.104a2.083 2.083 0 0 1-2.833.5L31.25 25"/></g>`), g.Group(children),
	)
}

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 43.75V6.25"/><path stroke="#344054" d="M31.25 12.5L25 6.25l-6.25 6.25"/></g>`), g.Group(children),
	)
}

func UpDirectionTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M37.854 36.896L25 24.042L12.146 36.896"/><path stroke="#306CFE" d="M12.146 36.896L6.25 31.02l17.27-17.292a2.08 2.08 0 0 1 2.96 0L43.75 31l-5.896 5.896"/></g>`), g.Group(children),
	)
}

func UpDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 41.667V8.333m6.25 6.25l-6.25-6.25l-6.25 6.25"/><path stroke="#306CFE" d="m27.083 35.417l6.25 6.25l6.25-6.25m-6.25-27.084v33.334"/></g>`), g.Group(children),
	)
}

func UpDownArrowTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M25 6.25v37.5"/><path stroke="#344054" d="M18.75 12.5L25 6.25l6.25 6.25m0 25L25 43.75l-6.25-6.25"/></g>`), g.Group(children),
	)
}

func UpJunctionSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 43.75h16.666m8.334-8.333h2.083zm0 8.333h2.083zM8.333 35.417H6.25zm0 8.333H6.25z"/><path stroke="#306CFE" d="M33.333 35.417h-4.166V18.75h6.25L26.604 8.167a2.083 2.083 0 0 0-3.208 0L14.583 18.75h6.25v16.667h-4.166"/></g>`), g.Group(children),
	)
}

func UpOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 33.333V16.667m-5.208 5.208L25 16.667l5.208 5.208"/><path stroke="#306CFE" d="M43.75 17.23v15.54L32.77 43.75H17.23L6.25 32.77V17.23L17.23 6.25h15.54z"/></g>`), g.Group(children),
	)
}

func UpTrend(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 12.5L29.167 27.083l-6.25-6.25L6.25 37.5"/><path stroke="#344054" d="M43.75 20.833V12.5h-8.333"/></g>`), g.Group(children),
	)
}

func Update(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M8.333 25a16.667 16.667 0 0 1 31.105-8.333M41.667 25a16.667 16.667 0 0 1-31.104 8.333"/><path stroke="#344054" d="M29.167 16.667h10.416V6.25m-18.75 27.083H10.417V43.75"/></g>`), g.Group(children),
	)
}

func UploadAltFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 22.917V6.25m6.25 6.25L25 6.25l-6.25 6.25"/><path stroke="#306CFE" d="M43.75 27.083L37.5 8.333M6.25 27.083l6.25-18.75m31.25 18.75v14.584a2.083 2.083 0 0 1-2.083 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083V27.083h10.417a8.333 8.333 0 1 0 16.666 0z"/></g>`), g.Group(children),
	)
}

func UploadFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 12.5L25 6.25l6.25 6.25M25 6.25v29.167"/><path stroke="#306CFE" d="M41.667 35.417v6.25a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083v-6.25"/></g>`), g.Group(children),
	)
}

func UploadNew(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 25v18.75m6.479-14.625L24.854 22.5l-6.334 6.333"/><path stroke="#306CFE" d="M12.5 31.48a13.73 13.73 0 1 1 20.833-15.084a9.6 9.6 0 0 1 1.73-.167A8.73 8.73 0 0 1 37.5 33.333"/></g>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M18.75 31.25h12.5a10.417 10.417 0 0 1 10.417 10.417a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083A10.417 10.417 0 0 1 18.75 31.25"/><path stroke="#306CFE" d="M25 22.917A8.333 8.333 0 1 0 25 6.25a8.333 8.333 0 0 0 0 16.667"/></g>`), g.Group(children),
	)
}

func UserAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M37.604 42.708h-.208"/><path stroke="#344054" stroke-width="2" d="M37.5 27.083v6.25"/><path stroke="#306CFE" stroke-width="2" d="M27.083 43.75H12.5a2.083 2.083 0 0 1-2.083-2.083a14.59 14.59 0 0 1 6.375-12.021l.166-.104a2.084 2.084 0 0 0 .417-2.917a12.35 12.35 0 0 1-2.792-7.875a12.5 12.5 0 0 1 24.834-2.083c.12.687.176 1.385.166 2.083"/></g>`), g.Group(children),
	)
}

func UserCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 31.25a8.333 8.333 0 1 0 0-16.667a8.333 8.333 0 0 0 0 16.667"/><path stroke="#344054" d="M29.52 29.896a8.23 8.23 0 0 1-9.04 0a14.58 14.58 0 0 0-8.917 8.166a18.75 18.75 0 0 0 26.687.188l.188-.188a14.6 14.6 0 0 0-8.917-8.166"/><path stroke="#306CFE" d="M43.75 25c0-10.355-8.395-18.75-18.75-18.75S6.25 14.645 6.25 25S14.645 43.75 25 43.75S43.75 35.355 43.75 25"/></g>`), g.Group(children),
	)
}

func UserSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 28.542a14.58 14.58 0 0 0-8.334 13.125a2.083 2.083 0 0 0 2.084 2.083h29.166a2.084 2.084 0 0 0 2.084-2.083a14.58 14.58 0 0 0-8.334-13.188"/><path stroke="#306CFE" d="M25 31.25c6.904 0 12.5-5.596 12.5-12.5S31.904 6.25 25 6.25s-12.5 5.596-12.5 12.5s5.596 12.5 12.5 12.5"/></g>`), g.Group(children),
	)
}

func UserSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 27.083h4.166a14.583 14.583 0 0 1 14.584 14.584a2.083 2.083 0 0 1-2.084 2.083H10.417a2.083 2.083 0 0 1-2.084-2.083a14.583 14.583 0 0 1 14.584-14.584"/><path stroke="#306CFE" d="M25 27.083c5.753 0 10.417-4.663 10.417-10.416S30.753 6.25 25 6.25s-10.417 4.664-10.417 10.417S19.247 27.083 25 27.083"/></g>`), g.Group(children),
	)
}

func VaseTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M32.063 21.52a10.2 10.2 0 0 1 3.354 7.25a48.7 48.7 0 0 1-1.646 13.418a2.08 2.08 0 0 1-2.084 1.562H18.272a2.08 2.08 0 0 1-2.084-1.562a48.7 48.7 0 0 1-1.604-13.417a10.2 10.2 0 0 1 3.334-7.25a8.33 8.33 0 0 0 2.916-6.125V6.25h8.334v9.146a8.33 8.33 0 0 0 2.895 6.125"/><path stroke="#344054" d="M31.25 6.25h-12.5"/></g>`), g.Group(children),
	)
}

func VestTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M21.042 39.583h7.916m3.792-31.25h-15.5"/><path stroke="#306CFE" d="M43.75 20.833v20.834a2.083 2.083 0 0 1-2.083 2.083h-12.5V27.083L33 7.917a2.08 2.08 0 0 1 2.083-1.667H37.5a2.083 2.083 0 0 1 2.083 2.083v8.334a4.167 4.167 0 0 0 4.167 4.166m-33.333-4.166V8.333A2.083 2.083 0 0 1 12.5 6.25h2.458a2.084 2.084 0 0 1 2.084 1.667l3.791 19.166V43.75h-12.5a2.083 2.083 0 0 1-2.083-2.083V20.833a4.167 4.167 0 0 0 4.167-4.166"/></g>`), g.Group(children),
	)
}

func Vibrate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M6.25 16.667v16.666m37.5-16.666v16.666z"/><path stroke="#306CFE" d="M33.333 41.667H16.667a2.083 2.083 0 0 1-2.084-2.084V10.417a2.083 2.083 0 0 1 2.084-2.084h16.666a2.083 2.083 0 0 1 2.084 2.084v29.166a2.083 2.083 0 0 1-2.084 2.084M29.167 8.333h-8.334l1.042 4.167h6.25z"/></g>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m35.417 30.458l5.312 2.646a2.083 2.083 0 0 0 3.021-1.854v-12.5a2.084 2.084 0 0 0-3.02-1.854l-5.313 2.645zM22.917 25l-4.167-3.125v6.25z"/><path stroke="#306CFE" d="M33.333 12.5h-25c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.933 2.083 2.083 2.083h25c1.15 0 2.084-.933 2.084-2.083V14.583c0-1.15-.933-2.083-2.084-2.083"/></g>`), g.Group(children),
	)
}

func Vision(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 29.167a4.167 4.167 0 1 0 0-8.334a4.167 4.167 0 0 0 0 8.334"/><path stroke="#306CFE" d="M43.75 25S37.5 37.5 25 37.5S6.25 25 6.25 25S12.5 12.5 25 12.5S43.75 25 43.75 25"/></g>`), g.Group(children),
	)
}

func Voice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="#306CFE" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.25 25h2.083a4.167 4.167 0 0 1 4.167 4.167v2.083a4.167 4.167 0 1 0 8.333 0v-12.5a4.167 4.167 0 0 1 8.334 0v12.5a4.167 4.167 0 1 0 8.333 0v-2.083A4.167 4.167 0 0 1 41.667 25h2.083"/>`), g.Group(children),
	)
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 8.333H10.417A4.167 4.167 0 0 0 6.25 12.5v0a4.167 4.167 0 0 0 4.167 4.167H43.75"/><path stroke="#306CFE" d="M43.75 16.667v22.916a2.083 2.083 0 0 1-2.083 2.084H8.333a2.083 2.083 0 0 1-2.083-2.084V12.5a4.167 4.167 0 0 0 4.167 4.167z"/><path stroke="#344054" d="M33.333 25H43.75v8.333H33.333a2.083 2.083 0 0 1-2.083-2.083v-4.167A2.083 2.083 0 0 1 33.333 25"/></g>`), g.Group(children),
	)
}

func WalletAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m12.5 14.583l16.938-7.52a2.083 2.083 0 0 1 2.708.958l3.27 6.562"/><path stroke="#306CFE" d="M41.667 33.333v8.334a2.083 2.083 0 0 1-2.084 2.083H8.333a2.083 2.083 0 0 1-2.083-2.083v-25a2.083 2.083 0 0 1 2.083-2.084h31.25a2.083 2.083 0 0 1 2.084 2.084V25"/><path stroke="#344054" d="M41.667 25h-6.25c-1.15 0-2.084.933-2.084 2.083v4.167c0 1.15.933 2.083 2.084 2.083h6.25c1.15 0 2.083-.932 2.083-2.083v-4.167c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func WalletAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M41.667 33.333v8.334a2.083 2.083 0 0 1-2.084 2.083H10.417a4.167 4.167 0 0 1-4.167-4.167V11.208"/><path stroke="#306CFE" d="M41.667 25v-8.333a2.083 2.083 0 0 0-2.084-2.084H10.417A4.167 4.167 0 0 1 6.25 9.708a4.31 4.31 0 0 1 4.396-3.458h22.687a2.083 2.083 0 0 1 2.084 2.083v6.25"/><path stroke="#344054" d="M41.667 25h-6.25c-1.15 0-2.084.933-2.084 2.083v4.167c0 1.15.933 2.083 2.084 2.083h6.25c1.15 0 2.083-.932 2.083-2.083v-4.167c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func WalletMoney(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m35.854 18.75l3.73-5.333l-10.23-7.167l-8.77 12.5zM18.938 8.333L11.625 18.75h8.958L25 12.5z"/><path stroke="#306CFE" d="M41.667 18.75H8.333c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V20.833c0-1.15-.933-2.083-2.083-2.083"/><path stroke="#344054" d="M33.333 27.083H43.75v8.334H33.333a2.083 2.083 0 0 1-2.083-2.084v-4.166a2.084 2.084 0 0 1 2.083-2.084"/></g>`), g.Group(children),
	)
}

func WalletMoneyThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="3" d="M35.313 31.25h.208"/><path stroke="#344054" stroke-width="2" d="m35.854 18.75l3.73-5.333l-10.23-7.167l-8.77 12.5zM18.938 6.25l-8.771 12.5h10.416l5.334-7.604z"/><path stroke="#306CFE" stroke-width="2" d="M41.667 18.75H8.333c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V20.833c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func WarningAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M25 18.75v8.333"/><path stroke="#344054" stroke-width="3" d="M25.104 35.417h-.208"/><path stroke="#306CFE" stroke-width="2" d="M21.354 8.73L5.48 37.5a4.166 4.166 0 0 0 3.646 6.25h31.75a4.167 4.167 0 0 0 3.646-6.25L28.646 8.73a4.166 4.166 0 0 0-7.292 0"/></g>`), g.Group(children),
	)
}

func WarningAltThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="#344054" stroke-width="2" d="M25 14.583v12.5"/><path stroke="#344054" stroke-width="3" d="M25.104 35.417h-.208"/><path stroke="#306CFE" stroke-width="2" d="m36.417 7.292l9.625 16.666a2.08 2.08 0 0 1 0 2.084l-9.625 16.666a2.08 2.08 0 0 1-1.792 1.042h-19.25a2.08 2.08 0 0 1-1.792-1.042L3.958 26.042a2.08 2.08 0 0 1 0-2.084l9.625-16.666a2.08 2.08 0 0 1 1.792-1.042h19.25a2.08 2.08 0 0 1 1.792 1.042"/></g>`), g.Group(children),
	)
}

func WashbasinFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 27.083h-4.166m-8.334-8.333v-8.333a4.167 4.167 0 1 1 8.334 0"/><path stroke="#306CFE" d="M8.333 43.75h33.334v-25H8.333zm-2.083-25h37.5zm2.083 25h12.5v-25h-12.5z"/></g>`), g.Group(children),
	)
}

func WaterBottle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 10.417h10.417a4.167 4.167 0 0 1 4.166 4.166V25a4.167 4.167 0 0 1-4.166 4.167H31.25"/><path stroke="#306CFE" d="M25 8.333a2.083 2.083 0 0 0-2.083-2.083H18.75a2.083 2.083 0 0 0-2.083 2.083v6.25H25zm2.083 6.25h-12.5a4.167 4.167 0 0 0-4.166 4.167v22.917A2.083 2.083 0 0 0 12.5 43.75h16.667a2.083 2.083 0 0 0 2.083-2.083V18.75a4.167 4.167 0 0 0-4.167-4.167"/></g>`), g.Group(children),
	)
}

func WaterCanTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="m32.333 24.5l7.605-4.167l3.812 3.813l-10.5 10.5m-20.02-7.688a8.334 8.334 0 1 1 9.312-10.687"/><path stroke="#306CFE" d="M14.77 39.583h16.293a2.083 2.083 0 0 0 2.083-2.27l-1.709-18.75a2.083 2.083 0 0 0-2.083-1.896H16.48a2.083 2.083 0 0 0-2.083 1.896l-1.709 18.75a2.08 2.08 0 0 0 2.084 2.27"/></g>`), g.Group(children),
	)
}

func WaterTap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M22.917 10.417v8m-6.25-8h12.5z"/><path stroke="#306CFE" d="M6.25 18.75v16.667m29.167-12.5h-4.959a2.08 2.08 0 0 1-1.875-1.146l-.916-1.875a2.08 2.08 0 0 0-1.875-1.146h-5.75a2.08 2.08 0 0 0-1.875 1.146l-.917 1.875a2.08 2.08 0 0 1-1.875 1.146H6.25v8.333h27.083a2.083 2.083 0 0 1 2.084 2.083V37.5a2.083 2.083 0 0 0 2.083 2.083h4.167A2.083 2.083 0 0 0 43.75 37.5v-6.25a8.333 8.333 0 0 0-8.333-8.333"/></g>`), g.Group(children),
	)
}

func WhiteBoard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M16.667 31.25L12.5 43.75m20.833-12.5l4.167 12.5z"/><path stroke="#306CFE" d="M41.667 6.25H8.333c-1.15 0-2.083.933-2.083 2.083v20.834c0 1.15.933 2.083 2.083 2.083h33.334c1.15 0 2.083-.933 2.083-2.083V8.333c0-1.15-.933-2.083-2.083-2.083"/></g>`), g.Group(children),
	)
}

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M20.833 37.5a4.167 4.167 0 1 0 8.334 0a4.167 4.167 0 0 0-8.334 0"/><path stroke="#306CFE" d="M6.25 15.208a29.17 29.17 0 0 1 37.5 0"/><path stroke="#306CFE" d="M38.438 21.583a20.83 20.83 0 0 0-26.876 0m21.501 6.355a12.5 12.5 0 0 0-16.126 0"/></g>`), g.Group(children),
	)
}

func Window(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M39.583 25H10.417M25 6.25v37.5z"/><path stroke="#306CFE" d="M6.25 43.75h37.5zm33.333 0H10.417V8.333A2.083 2.083 0 0 1 12.5 6.25h25a2.083 2.083 0 0 1 2.083 2.083z"/></g>`), g.Group(children),
	)
}

func WindowFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M25 39.583H10.417V6.25H22.02"/><path stroke="#306CFE" d="M30.708 27.646C25 32.687 25 43.75 25 43.75h12.5a2.083 2.083 0 0 0 2.083-2.083V8.333A2.083 2.083 0 0 0 37.5 6.25H22.02a28.895 28.895 0 0 0 17.563 27.083"/></g>`), g.Group(children),
	)
}

func WindowSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#344054" d="M33.333 22.917v8.333zm8.334-8.334H8.333v8.334h33.334z"/><path stroke="#306CFE" d="M6.25 6.25h37.5m-2.083 0H8.333v37.5h33.334zM6.25 43.75h37.5z"/></g>`), g.Group(children),
	)
}

func Work(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M33.333 8.333A2.083 2.083 0 0 0 31.25 6.25h-12.5a2.083 2.083 0 0 0-2.083 2.083v6.25h16.666zM43.75 41.667v-25a2.083 2.083 0 0 0-2.083-2.084H8.333a2.083 2.083 0 0 0-2.083 2.084v25a2.083 2.083 0 0 0 2.083 2.083h33.334a2.083 2.083 0 0 0 2.083-2.083"/><path stroke="#344054" d="M22.917 29.167H18a8.33 8.33 0 0 1-7.583-5.042l-3.792-8.646a2.08 2.08 0 0 1 1.708-.896h33.334a2.08 2.08 0 0 1 1.708.896l-3.792 8.646A8.33 8.33 0 0 1 32 29.167h-4.917"/><path stroke="#306CFE" d="M27.083 27.083h-4.166v4.167h4.166z"/></g>`), g.Group(children),
	)
}

func WorkAgenda(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M6.25 33.333h4.167m-2.084 8.334V8.333a2.083 2.083 0 0 1 2.084-2.083h31.25a2.083 2.083 0 0 1 2.083 2.083v33.334a2.083 2.083 0 0 1-2.083 2.083h-31.25a2.083 2.083 0 0 1-2.084-2.083m-2.083-25h4.167zm0 8.333h4.167z"/><path stroke="#344054" d="M41.667 43.75h-6.25V6.25h6.25a2.083 2.083 0 0 1 2.083 2.083v33.334a2.083 2.083 0 0 1-2.083 2.083"/></g>`), g.Group(children),
	)
}

func ZigZagLeftRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M43.75 37.5h-25a6.25 6.25 0 0 1-6.25-6.25v0A6.25 6.25 0 0 1 18.75 25h12.5a6.25 6.25 0 0 0 0-12.5h-25"/><path stroke="#344054" d="M39.583 41.667L43.75 37.5l-4.167-4.167m-29.166-25L6.25 12.5l4.167 4.167"/></g>`), g.Group(children),
	)
}

func ZigZagLeftUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M39.583 43.75v-25a6.25 6.25 0 1 0-12.5 0V37.5a6.25 6.25 0 1 1-12.5 0V6.25"/><path stroke="#344054" d="M18.75 10.417L14.583 6.25l-4.166 4.167"/></g>`), g.Group(children),
	)
}

func ZigZagRightUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke="#306CFE" d="M35.417 6.25V37.5a6.25 6.25 0 0 1-6.25 6.25v0a6.25 6.25 0 0 1-6.25-6.25V18.75a6.25 6.25 0 0 0-12.5 0v25"/><path stroke="#344054" d="M39.583 10.417L35.417 6.25l-4.167 4.167"/></g>`), g.Group(children),
	)
}
