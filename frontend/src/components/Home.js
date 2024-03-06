import React from 'react'

function Home() {
  return (
    <div className="homePage">
            <header className="heroSection">
                <h1>Welcome to WebMarket!</h1>
                <p>Your one-stop shop for everything.</p>
            </header>
            <section className="featuresSection">
                <div className="feature">
                    <h2>Wide Selection</h2>
                    <p>Explore a wide range of products across various categories.</p>
                </div>
                <div className="feature">
                    <h2>Top Deals</h2>
                    <p>Discover top deals and offers exclusively available for you.</p>
                </div>
                <div className="feature">
                    <h2>24/7 Support</h2>
                    <p>Our team is here to help you anytime, anywhere.</p>
                </div>
            </section>
        </div>
  )
}

export default Home