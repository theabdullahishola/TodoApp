import logo from "../assets/Go-Logo_LightBlue.png"
import image2 from "../assets/react.png" 
import emoji from "../assets/emoji.png" 

export default function Navbar() {
  return (
    <>
      <div className="image-container">
            <div className="img-ctn">
             <img src={logo}  alt="Go Logo" />
        
                <img src={image2} alt="Another Image" />

             <img src={emoji} alt="bursted emoji" />
        </div>
        <div className="toggle">
            <label className="switch">
            <input type="checkbox"/>
            <span className="slider"></span>
            
            </label>
        </div>
        
      </div>
    </>
  )
}