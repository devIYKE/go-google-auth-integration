// Main JavaScript file for the application

document.addEventListener('DOMContentLoaded', function() {
    // Set current year in the footer
    const yearElement = document.querySelector('footer p');
    if (yearElement) {
        const currentYear = new Date().getFullYear();
        yearElement.innerHTML = yearElement.innerHTML.replace('{{.CurrentYear}}', currentYear);
    }

    // Add animation to the Google Sign-In button
    const googleBtn = document.querySelector('.google-btn');
    if (googleBtn) {
        googleBtn.addEventListener('mouseover', function() {
            this.style.transform = 'scale(1.03)';
        });
        googleBtn.addEventListener('mouseout', function() {
            this.style.transform = 'scale(1)';
        });
    }
});
