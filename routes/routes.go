package routes

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
)

const baseURL = "http://10.49.122.144:1208"

const userSecret = "9759b22e3ea6d95d2e0606132358387e5834a8682dd28ee7ebea027da9d76440"

func SetupRoutes(app *fiber.App) {
    app.Post("/signup", signup)
    app.Post("/check", check)
    app.Post("/getUserSecret", getUserSecret)
    app.Post("/getUserLevel", getUserLevel)
    app.Post("/getUserPoints", getUserPoints)
    app.Post("/iNeedAHint", iNeedAHint)
    app.Post("/enterChallenge", enterChallenge)
    app.Post("/submitSolution", submitSolution)
}

func signup(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s/signup", baseURL)
	bodyReader := bytes.NewReader(c.Request().Body())
	resp, err := http.Post(url, "application/json", bodyReader)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.Status(resp.StatusCode).Send(body)
}

func check(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s/check", baseURL)
	bodyReader := bytes.NewReader(c.Request().Body())
	resp, err := http.Post(url, "application/json", bodyReader)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.Status(resp.StatusCode).Send(body)
}


func getUserSecret(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/getUserSecret", baseURL)
    return postRequest(c, url)
}

func getUserLevel(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/getUserLevel", baseURL)
    return postRequest(c, url)
}

func getUserPoints(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/getUserPoints", baseURL)
    return postRequest(c, url)
}

func iNeedAHint(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/iNeedAHint", baseURL)
    return postRequest(c, url)
}

func enterChallenge(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/enterChallenge", baseURL)
    return postRequest(c, url)
}

func submitSolution(c *fiber.Ctx) error {
    url := fmt.Sprintf("%s/submitSolution", baseURL)
    return postRequest(c, url)
}

func postRequest(c *fiber.Ctx, url string) error {
    bodyReader := bytes.NewReader(c.Request().Body())
    resp, err := http.Post(url, "application/json", bodyReader)
    if err != nil {
        return c.Status(500).SendString("Internal Server Error")
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return c.Status(500).SendString("Internal Server Error")
    }

    fmt.Println("Response:", string(body))

    return c.Status(resp.StatusCode).Send(body)
}