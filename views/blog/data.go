package blog

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Article struct {
	ID       int
	Title    string
	Author   string
	Date     time.Time
	Summary  string
	ImageUrl string
	Category string
	Content  string
}

var AllArticles = []Article{
	{ID: 1, Title: "The Future of AI in Healthcare", Author: "Dr. Jane Smith", Date: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), Summary: "Exploring how artificial intelligence is revolutionizing medical diagnostics and treatment plans.", ImageUrl: "/images/ai-healthcare.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 2, Title: "Sustainable Travel: Eco-Friendly Destinations", Author: "Alex Green", Date: time.Date(2024, 1, 18, 0, 0, 0, 0, time.UTC), Summary: "Discover top destinations that prioritize environmental conservation and sustainable tourism.", ImageUrl: "/images/eco-travel.jpg", Category: "Travel", Content: "Full content here..."},
	{ID: 3, Title: "The Rise of Plant-Based Diets", Author: "Nutrition Expert Sarah Lee", Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), Summary: "Examining the health benefits and environmental impact of adopting a plant-based lifestyle.", ImageUrl: "/images/plant-based-diet.jpg", Category: "Health", Content: "Full content here..."},
	{ID: 4, Title: "Cryptocurrency: A Beginner's Guide", Author: "Finance Guru Mike Ross", Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), Summary: "Understanding the basics of cryptocurrency and its potential impact on the future of finance.", ImageUrl: "/images/crypto-guide.jpg", Category: "Finance", Content: "Full content here..."},
	{ID: 5, Title: "Gourmet Cooking at Home: Chef's Secrets", Author: "Chef Michael Brown", Date: time.Date(2024, 1, 25, 0, 0, 0, 0, time.UTC), Summary: "Learn professional cooking techniques to elevate your home-cooked meals to restaurant quality.", ImageUrl: "/images/gourmet-cooking.jpg", Category: "Food", Content: "Full content here..."},
	{ID: 6, Title: "5G Technology: Transforming Connectivity", Author: "Tech Analyst Lisa Chen", Date: time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), Summary: "Exploring the capabilities of 5G networks and their potential to revolutionize various industries.", ImageUrl: "/images/5g-tech.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 7, Title: "Hidden Gems of Southeast Asia", Author: "Travel Blogger Tom Wilson", Date: time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC), Summary: "Uncovering lesser-known but breathtaking destinations across Southeast Asia.", ImageUrl: "/images/southeast-asia.jpg", Category: "Travel", Content: "Full content here..."},
	{ID: 8, Title: "Mindfulness Meditation: A Path to Mental Wellness", Author: "Psychologist Dr. Emily Carter", Date: time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC), Summary: "Understanding the science behind mindfulness and its benefits for mental health.", ImageUrl: "/images/mindfulness.jpg", Category: "Health", Content: "Full content here..."},
	{ID: 9, Title: "The Future of Electric Vehicles", Author: "Auto Industry Expert Jack Thompson", Date: time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC), Summary: "Analyzing the rapid advancements in electric vehicle technology and market trends.", ImageUrl: "/images/electric-vehicles.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 10, Title: "Investing in Renewable Energy", Author: "Financial Advisor Emma Rodriguez", Date: time.Date(2024, 2, 8, 0, 0, 0, 0, time.UTC), Summary: "Exploring investment opportunities in the growing renewable energy sector.", ImageUrl: "/images/renewable-energy.jpg", Category: "Finance", Content: "Full content here..."},
	{ID: 11, Title: "The Art of Sourdough Bread Making", Author: "Baker Paul White", Date: time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), Summary: "Mastering the techniques of creating perfect sourdough bread at home.", ImageUrl: "/images/sourdough-bread.jpg", Category: "Food", Content: "Full content here..."},
	{ID: 12, Title: "Space Tourism: The Next Frontier", Author: "Aerospace Engineer Dr. Samantha Lee", Date: time.Date(2024, 2, 13, 0, 0, 0, 0, time.UTC), Summary: "Examining the current state and future prospects of commercial space travel.", ImageUrl: "/images/space-tourism.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 13, Title: "Underwater Photography Tips", Author: "Marine Photographer David Ocean", Date: time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), Summary: "Expert advice on capturing stunning underwater images while diving or snorkeling.", ImageUrl: "/images/underwater-photo.jpg", Category: "Travel", Content: "Full content here..."},
	{ID: 14, Title: "The Science of Sleep", Author: "Sleep Researcher Dr. Mark Dreamer", Date: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), Summary: "Understanding the importance of sleep and strategies for improving sleep quality.", ImageUrl: "/images/sleep-science.jpg", Category: "Health", Content: "Full content here..."},
	{ID: 15, Title: "Blockchain Beyond Cryptocurrency", Author: "Tech Innovator Rachel Blockchain", Date: time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), Summary: "Exploring diverse applications of blockchain technology across various industries.", ImageUrl: "/images/blockchain-tech.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 16, Title: "Sustainable Fashion: Eco-Friendly Trends", Author: "Fashion Designer Olivia Green", Date: time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC), Summary: "Discovering the latest trends in sustainable and ethical fashion.", ImageUrl: "/images/sustainable-fashion.jpg", Category: "Lifestyle", Content: "Full content here..."},
	{ID: 17, Title: "Artificial Intelligence in Education", Author: "EdTech Specialist Prof. Alan Smart", Date: time.Date(2024, 2, 25, 0, 0, 0, 0, time.UTC), Summary: "Examining how AI is transforming learning experiences and educational systems.", ImageUrl: "/images/ai-education.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 18, Title: "Urban Farming: Growing Food in Cities", Author: "Urban Agriculturist Maria Gardens", Date: time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC), Summary: "Exploring innovative methods for cultivating fresh produce in urban environments.", ImageUrl: "/images/urban-farming.jpg", Category: "Food", Content: "Full content here..."},
	{ID: 19, Title: "The Psychology of Decision Making", Author: "Behavioral Economist Dr. Chris Choice", Date: time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC), Summary: "Understanding the cognitive processes behind our everyday choices and decisions.", ImageUrl: "/images/decision-making.jpg", Category: "Psychology", Content: "Full content here..."},
	{ID: 20, Title: "Remote Work Revolution", Author: "Workplace Analyst Jennifer Remote", Date: time.Date(2024, 3, 5, 0, 0, 0, 0, time.UTC), Summary: "Analyzing the long-term impacts of remote work on businesses and employees.", ImageUrl: "/images/remote-work.jpg", Category: "Business", Content: "Full content here..."},
	{ID: 21, Title: "The Future of Quantum Computing", Author: "Quantum Physicist Dr. Schrödinger Cat", Date: time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC), Summary: "Exploring the potential of quantum computers to revolutionize data processing and problem-solving.", ImageUrl: "/images/quantum-computing.jpg", Category: "Technology", Content: "Full content here..."},
	{ID: 22, Title: "Culinary Tourism: Eating Your Way Around the World", Author: "Food Critic Gourmet Gary", Date: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), Summary: "Discovering unique culinary experiences and food cultures across different countries.", ImageUrl: "/images/culinary-tourism.jpg", Category: "Travel", Content: "Full content here..."},
	{ID: 23, Title: "Biohacking: Optimizing Human Performance", Author: "Biohacker Zoe Upgrade", Date: time.Date(2024, 3, 13, 0, 0, 0, 0, time.UTC), Summary: "Exploring cutting-edge techniques for enhancing physical and cognitive abilities.", ImageUrl: "/images/biohacking.jpg", Category: "Health", Content: "Full content here..."},
	{ID: 24, Title: "The Rise of Decentralized Finance (DeFi)", Author: "Fintech Expert Lucas Blockchain", Date: time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), Summary: "Understanding the potential of DeFi to reshape traditional financial systems.", ImageUrl: "/images/defi.jpg", Category: "Finance", Content: "Full content here..."},
	{ID: 25, Title: "Vertical Gardening: Maximizing Small Spaces", Author: "Urban Gardener Lily Greenthumb", Date: time.Date(2024, 3, 18, 0, 0, 0, 0, time.UTC), Summary: "Innovative techniques for creating lush, productive gardens in limited urban spaces.", ImageUrl: "/images/vertical-gardening.jpg", Category: "Lifestyle", Content: "Full content here..."},
	// Add more articles as needed...
}

func GetRandomArticles(n int) []Article {
	shuffled := make([]Article, len(AllArticles))
	copy(shuffled, AllArticles)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	if n > len(shuffled) {
		n = len(shuffled)
	}
	return shuffled[:n]
}

func SearchArticles(query string, category string) []Article {
	results := make([]Article, len(AllArticles))
	copy(results, AllArticles)

	// Filter by category if specified
	if category != "" {
		filteredResults := []Article{}
		for _, article := range results {
			if strings.EqualFold(article.Category, category) {
				filteredResults = append(filteredResults, article)
			}
		}
		results = filteredResults
	}

	// If there's no query, just return the (possibly category-filtered) results
	if query == "" {
		return results
	}

	// Sort results by relevance
	sort.Slice(results, func(i, j int) bool {
		scoreI := relevanceScore(results[i], query)
		scoreJ := relevanceScore(results[j], query)
		return scoreI > scoreJ
	})

	return results
}

func relevanceScore(article Article, query string) int {
	score := 0
	lowercaseQuery := strings.ToLower(query)
	lowercaseTitle := strings.ToLower(article.Title)
	lowercaseSummary := strings.ToLower(article.Summary)
	lowercaseContent := strings.ToLower(article.Content)

	// Exact match in title
	if strings.Contains(lowercaseTitle, lowercaseQuery) {
		score += 100
	}

	// Word match in title
	for _, word := range strings.Fields(lowercaseQuery) {
		if strings.Contains(lowercaseTitle, word) {
			score += 50
		}
	}

	// Exact match in summary
	if strings.Contains(lowercaseSummary, lowercaseQuery) {
		score += 30
	}

	// Word match in summary
	for _, word := range strings.Fields(lowercaseQuery) {
		if strings.Contains(lowercaseSummary, word) {
			score += 15
		}
	}

	// Exact match in content
	if strings.Contains(lowercaseContent, lowercaseQuery) {
		score += 10
	}

	// Word match in content
	for _, word := range strings.Fields(lowercaseQuery) {
		if strings.Contains(lowercaseContent, word) {
			score += 5
		}
	}

	// Boost score for more recent articles
	daysOld := int(time.Since(article.Date).Hours() / 24)
	score += max(0, 100-daysOld) // Boost decreases as the article gets older, up to 100 days

	return score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
