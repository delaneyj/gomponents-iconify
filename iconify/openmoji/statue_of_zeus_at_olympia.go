package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatueOfZeusAtOlympia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="44.332" cy="11.416" r="2.871" fill="#fcea2b"/><path fill="#fcea2b" d="M37.917 33.173H25.323c-1.461 0-1.856-1.59-1.856-2.775c0-.553 2.682-.98 3.235-.98h12.287c.553 0-.072 2.201-.072 2.754a1 1 0 0 1-1 1Z"/><path fill="#fcea2b" d="M48.49 21.976c-.101-2.875-1.802-4.819-4.368-4.954a5.37 5.37 0 0 0-5.4 4.487l-.89 5.986v.003l-1.128 7.753l-4.257.001l-2.604.14a8.466 8.466 0 0 0-3.491 1.173a5.318 5.318 0 0 0-2.689 4.843l.332 15.942c0 1.302 0 4.014 2.414 4.014a3.228 3.228 0 0 0 3.403-2.651l1.853-16.009l13.058-.025a3.797 3.797 0 0 0 3.768-3.705Z"/><path fill="none" stroke="#9b9b9a" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M47.058 45.209v15.749M34.549 45.209v15.749"/><path fill="#9b9b9a" d="m47.44 25.417l3.68-.239l-.591 19.86l-19.187.171v-3.387l15.689-1.154l.409-15.251z"/><g fill="none" stroke="#000" stroke-width="2"><circle cx="44.332" cy="11.416" r="2.871" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m43.551 22.011l-2.09 12.971l-8.96-8.853"/><path stroke-linecap="round" stroke-linejoin="round" d="M31.882 41.675a1.107 1.107 0 0 0-1.105 1.104v1.326a1.188 1.188 0 0 0 1.105 1.104h18.113a1.188 1.188 0 0 0 1.105-1.104V25.86a1.107 1.107 0 0 0-1.105-1.105l-1.86.021c-.552 0-.462.477-.462 1.14l-.182 15.746Zm.168-11.277h-7.582m24.395 14.811v15.749M32.694 45.209v15.395m-9.89-33.712l1.793-5.708l1.865 5.708h-3.658zm-2.406-8.01a5.312 5.312 0 0 0 4.2 2.302s3.163 0 4.27-2.302m-9.15 44.704v-2.628h32.564v2.628M34.15 15.273l6.452 45.685"/><path stroke-linecap="round" stroke-linejoin="round" d="m38.82 27.643l.891-5.986a4.39 4.39 0 0 1 4.395-3.636h0c2.161.115 3.307 1.826 3.384 3.99l.001 16.963a2.77 2.77 0 0 1-2.769 2.705h-4.817l-9.132.028l-1.955 16.89c-.223 1.373-.926 2.361-2.041 2.361a2.221 2.221 0 0 1-2.144-2.498V41.345c0-2.982 2.41-5.094 5.797-5.094h2.07m-7.867-12.883v2.856"/><circle cx="24.633" cy="17.225" r=".653" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}