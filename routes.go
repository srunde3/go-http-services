package main

func (s *server) routes() {
	s.router.Get("/health", handleHealthCheck)
	s.router.Get("/hi-mom", handleHiMom)
	s.router.Get("/echo/{parameter}", handleEchoUrlParameter)
}
