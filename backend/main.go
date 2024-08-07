package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Resume struct {
    Name       string   `json:"name"`
    Email      string   `json:"email"`
    Interests  []string `json:"interests"`
    Languages  []string `json:"languages"`
    Skills     []string `json:"skills"`
    Licenses   []string `json:"licenses"`
    Education  []string `json:"education"`
    Activity   []string `json:"activity"`
}

func main() {
    r := gin.Default()
    r.Static("/static", "./frontend/build/static")
    r.StaticFile("/", "./frontend/build/index.html")
    r.GET("/api/resume", getResume)
    r.POST("/api/resume", createResume)
    r.Run(":8080")
}

func getResume(c *gin.Context) {
    data, err := ioutil.ReadFile("resume.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read resume file"})
        return
    }

    var resume Resume
    if err := json.Unmarshal(data, &resume); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse resume file"})
        return
    }

    c.JSON(http.StatusOK, resume)
}

func createResume(c *gin.Context) {
    var resume Resume
    if err := c.ShouldBindJSON(&resume); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    data, err := json.Marshal(resume)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize resume"})
        return
    }

    if err := ioutil.WriteFile("resume.json", data, 0644); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write resume file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Resume created"})
}
